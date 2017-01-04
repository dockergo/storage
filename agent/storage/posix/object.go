package posix

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/flyaways/storage/agent/constant"
	"github.com/flyaways/storage/agent/protocol"
	"github.com/flyaways/storage/agent/result"
	"github.com/flyaways/storage/agent/util"
	"github.com/flyaways/storage/agent/util/log"
)

func (posix *Posix) PutObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	rawRequestdata, finalkey, err := protocol.PutHeadchecker(ctx, res, bucket, key)
	if err != nil {
		return
	}

	posix.uploadObject(ctx, res, rawRequestdata, finalkey, bucket)
	return
}

func (posix *Posix) PostObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParamPost(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	rawRequestdata, finalkey, err := protocol.PostHeadchecker(ctx, res, bucket, key)
	if err != nil {
		return
	}

	posix.uploadObject(ctx, res, rawRequestdata, finalkey, bucket)
	return
}

func (posix *Posix) uploadObject(ctx *gin.Context, res *result.Result, rawRequestdata []byte, finalkey, bucket string) {
	err := os.MkdirAll(posix.getBucketPath(bucket), os.ModePerm)
	if err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(err)
		return
	}

	if err = posix.Checker.DirChecker(posix.getBucketPath(bucket)); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(err)
		return
	}

	filename := filepath.Join(posix.getBucketPath(bucket), finalkey)
	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return
	}

	if err := ioutil.WriteFile(filename, rawRequestdata, 0666); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(err)
		return
	}

	ctx.Header(constant.ETag, util.GetETagValue(rawRequestdata))
	ctx.Header(constant.NewFileName, finalkey)
	ctx.Status(http.StatusOK)
	ctx.JSON(http.StatusOK, gin.H{constant.NewFileName: finalkey})

	return
}

func (posix *Posix) GetObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	if err := posix.Checker.DirChecker(posix.getBucketPath(bucket)); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(err)
		return
	}
	filename := filepath.Join(posix.getBucketPath(bucket), key)
	if err := posix.Checker.FileChecker(filename); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(err)
		return
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(err)
		return
	}

	protocol.GetCkecker(ctx, content, res)
	return
}

func (posix *Posix) HeadObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	if err := posix.Checker.DirChecker(posix.getBucketPath(bucket)); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(err)
		return
	}

	filename := filepath.Join(posix.getBucketPath(bucket), key)
	if err := posix.Checker.FileChecker(filename); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(err)
		return
	}

	fileinfo, err := os.Stat(filename)
	if err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(err)
		return
	}

	//no check
	objEtag := ctx.Request.Header.Get(constant.ETag)
	if protocol.CheckETag(ctx, objEtag) {
		log.Error("[%s:CheckETag]")
		res.Error(errors.New("invalid Etag"))
		return
	}

	lastModified := fileinfo.ModTime().Format(constant.TimeFormat)

	ctx.Header(constant.ContentLength, strconv.FormatInt(fileinfo.Size(), 10))
	if fileinfo.Size() <= 0 {
		ctx.Header(constant.ContentLength, "0")
		ctx.Header(constant.AcceptRanges, "bytes")
	}

	protocol.HeadChecker(ctx, res, objEtag, lastModified)
	return
}

func (posix *Posix) DeleteObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}
	if err := posix.Checker.DirChecker(posix.getBucketPath(bucket)); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(err)
		return
	}

	filename := filepath.Join(posix.getBucketPath(bucket), key)
	if err := posix.Checker.FileChecker(filename); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(err)
		return
	}

	err := os.Remove(filename)
	if err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
	return
}
