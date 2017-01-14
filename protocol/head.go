package protocol

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/flyaways/storage/constant"
	errs "github.com/flyaways/storage/errors"
	"github.com/flyaways/storage/result"
	"github.com/flyaways/storage/util"
	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
)

func PostHeader(ctx *gin.Context, res *result.Result, bucket, key string) ([]byte, string, error) {
	theReader, _, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Error("[object.POST read multipart error:%s]", err.Error())
		res.Error(errs.InternalError)
		return nil, "", err
	}

	data, err := ioutil.ReadAll(theReader)
	if err != nil {
		log.Error("[PostHeader:%s]", err.Error())
		res.Error(errs.InternalError)
		return nil, "", err
	}

	key = computeKey(ctx, key, data)
	return data, key, nil
}

func PutHeader(ctx *gin.Context, res *result.Result, bucket, key string) ([]byte, string, error) {
	if ctx.Request.Body == nil {
		log.Error("[object.PUT body is nil]")
		res.Error(errs.FileEmpty)
		return nil, "", errors.New("ReqBodyNil")
	}

	defer ctx.Request.Body.Close()
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Error("[PutHeader:%s]", err.Error())
		res.Error(errs.InternalError)
		return nil, "", err
	}

	key = computeKey(ctx, key, data)
	return data, key, nil
}

func GetHeader(ctx *gin.Context, content []byte, res *result.Result) {
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

func computeKey(ctx *gin.Context, key string, data []byte) string {
	return util.GetSha1Hex(data)
}
