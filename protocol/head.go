package protocol

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/flyaways/storage/constant"
	errs "github.com/flyaways/storage/errors"
	"github.com/flyaways/storage/result"
	"github.com/flyaways/storage/util"
	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
)

func PostHeadchecker(ctx *gin.Context, res *result.Result, bucket, key string) ([]byte, string, error) {
	theReader, _, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Error("[object.POST read multipart error:%s]", err.Error())
		res.Error(errs.InternalError)
		return nil, "", err
	}

	rawRequestdata, err := ioutil.ReadAll(theReader)
	if err != nil {
		log.Error("[object.POST read multipart to bytes error:%s]", err.Error())
		res.Error(errs.InternalError)
		return nil, "", err
	}

	contentLength := ctx.Request.Header.Get(constant.ContentLength)
	if contentLength == "" || contentLength == "0" {
		log.Error("[object.POST content length is empty]")
		res.Error(errs.MissingContentLength)
		return rawRequestdata, "", errors.New("MissingContentLength")
	}

	if ctx.Request.Header.Get(constant.ContentMD5) != "" &&
		ctx.Request.Header.Get(constant.ContentMD5) != util.GetBase64Md5(rawRequestdata) {
		log.Error("[object.POST MD5DoesNotMatch]")
		res.Error(errs.MD5DoesNotMatch)
		return rawRequestdata, "", errors.New("MD5DoesNotMatch")
	}

	finalkey := ComputeKey(ctx, key, rawRequestdata)
	return rawRequestdata, finalkey, nil
}

func PutHeadchecker(ctx *gin.Context, res *result.Result, bucket, key string) ([]byte, string, error) {
	if ctx.Request.Body == nil {
		log.Error("[object.PUT body is nil]")
		res.Error(errs.ReqBodyNil)
		return nil, "", errors.New("ReqBodyNil")
	}

	defer ctx.Request.Body.Close()
	rawRequestdata, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Error("[object.POST read multipart to bytes error:%s]", err.Error())
		//res.Error(errs.PermissionDenied)
		return nil, "", err
	}
	contentLength := ctx.Request.Header.Get(constant.ContentLength)
	if contentLength == "" || contentLength == "0" {
		log.Error("[object.PUT content length is empty]")
		res.Error(errs.MissingContentLength)
		return rawRequestdata, "", errors.New("MissingContentLength")
	}

	if ctx.Request.Header.Get(constant.ContentMD5) != "" &&
		ctx.Request.Header.Get(constant.ContentMD5) != util.GetBase64Md5(rawRequestdata) {
		log.Error("[MD5DoesNotMatch:%s]", err.Error())
		res.Error(errs.MD5DoesNotMatch)
		return rawRequestdata, "", errors.New("MD5DoesNotMatch")
	}
	finalkey := ComputeKey(ctx, key, rawRequestdata)
	return rawRequestdata, finalkey, nil
}

func GetCkecker(ctx *gin.Context, content []byte, res *result.Result) {
	objContentLength := ctx.Request.Header.Get(constant.ContentLength)
	if objContentLength == "" {
		objContentLength = "0"
	}

	sendSize, err := strconv.ParseInt(objContentLength, 10, 64)
	if err != nil {
		log.Error("[obj.GET object error:%s]", err.Error())
		res.Error(err)
		return
	}

	reqRange := ctx.Request.Header.Get(constant.Range)
	ranges, err := ParseRange(reqRange, sendSize)

	if err != nil {
		res.Error(errs.InvalidRange)
		log.Error("[obj.GET object error:%s]", err.Error())
		return
	}

	if SumRangesSize(ranges) > sendSize {
		ranges = nil
	}

	switch {
	case len(ranges) == 1:
		partRange := ranges[0]
		//content = content[partRange.start : partRange.start+partRange.length]
		sendlength := partRange.Length
		ctx.Header(constant.ContentRange, partRange.ContentRange(sendlength))
	}

	ctx.Writer.Write(content)

	responseMake(ctx)

	ctx.Status(http.StatusOK)
	return
}

func HeadChecker(ctx *gin.Context, res *result.Result, objEtag, lastModified string) {

	modtime, err := time.Parse(constant.TimeFormat, lastModified)
	if err != nil {
		log.Error("[obj.HEAD object LastModified error:%s]", err.Error())
		res.Error(err)
		return
	}

	done, err := CheckModSince(ctx, modtime)
	if err != nil || done {
		log.Error("[obj.HEAD object modSince error:%s]", err.Error())
		res.Error(err)
		return
	}

	ctx.Header(constant.LastModified, lastModified)
	ctx.Header(constant.Created, lastModified)
	ctx.Header(constant.ETag, objEtag)
	ctx.Status(http.StatusOK)
}
