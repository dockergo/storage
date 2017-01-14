package protocol

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"io"

	"github.com/flyaways/storage/constant"
	errs "github.com/flyaways/storage/errors"
	"github.com/flyaways/storage/result"
	"github.com/flyaways/storage/util"
	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
)

func Header(ctx *gin.Context, res *result.Result, bucket, key string) ([]byte, string, error) {
	var body io.Reader
	var err error
	if ctx.Request.Method == "POST" {
		body, _, err = ctx.Request.FormFile("file")
		if err != nil {
			log.Error("[object.POST read multipart error:%s]", err.Error())
			res.Error(errs.InternalError)
			return nil, "", err
		}
	} else {
		if ctx.Request.Body == nil {
			log.Error("[object.PUT body is nil]")
			res.Error(errs.FileEmpty)
			return nil, "", errors.New("[reqest body nil]")
		}
		body = ctx.Request.Body
		defer ctx.Request.Body.Close()
	}

	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Error("[Header:%s]", err.Error())
		res.Error(errs.InternalError)
		return nil, "", err
	}

	//key = KeyMaker(ctx, key, data)
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
		content = content[partRange.Start : partRange.Start+partRange.Length]
		sendlength := partRange.Length
		ctx.Header(constant.ContentRange, partRange.ContentRange(sendlength))
	}

	ctx.Writer.Write(content)

	responseMake(ctx)

	ctx.Status(http.StatusOK)
	return
}

func KeyMaker(ctx *gin.Context, key string, data []byte) string {
	return util.GetSha1Hex(data)
}
