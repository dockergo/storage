package posix

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"

	"fmt"

	"github.com/flyaways/storage/constant"
	"github.com/flyaways/storage/errors"
	"github.com/flyaways/storage/protocol"
	"github.com/flyaways/storage/result"
	"github.com/flyaways/storage/util"
	"github.com/flyaways/storage/util/log"
)

func (posix *Posix) PutObject(ctx *gin.Context) {
	res, bucket, key := protocol.Param(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	data, finalkey, err := protocol.Header(ctx, res, bucket, key)
	if err != nil {
		return
	}

	posix.uploadObject(ctx, res, data, finalkey, bucket)
	return
}

func (posix *Posix) PostObject(ctx *gin.Context) {
	res, bucket, key := protocol.Param(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	data, finalkey, err := protocol.Header(ctx, res, bucket, key)
	if err != nil {
		return
	}

	posix.uploadObject(ctx, res, data, finalkey, bucket)
	return
}

func (posix *Posix) uploadObject(ctx *gin.Context, res *result.Result, data []byte, finalkey, bucket string) {
	err := os.MkdirAll(posix.getBucketPath(bucket), os.ModePerm)
	if err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.InternalError)
		return
	}

	if err := posix.DirChecker(posix.getBucketPath(bucket)); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.NoSuchBucket)
		return
	}

	filename := filepath.Join(posix.getBucketPath(bucket), finalkey)
	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return
	}

	if err = ioutil.WriteFile(filename, data, 0666); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.NoSuchKey)
		return
	}

	ctx.Header(constant.ETag, util.GetETagValue(data))
	ctx.Header(constant.NewFileName, finalkey)
	ctx.Status(http.StatusOK)
	ctx.JSON(http.StatusOK, gin.H{constant.NewFileName: finalkey})

	return
}

func (posix *Posix) GetObject(ctx *gin.Context) {
	res, bucket, key := protocol.Param(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	if err := posix.DirChecker(posix.getBucketPath(bucket)); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.NoSuchBucket)
		return
	}

	filename := filepath.Join(posix.getBucketPath(bucket), key)
	if err := posix.FileChecker(filename); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.NoSuchKey)
		return
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.NoSuchKey)
		return
	}

	fileinfo, err := os.Stat(filename)
	if err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.NoSuchKey)
		return
	} else {
		ctx.Header(constant.ContentLength, strconv.FormatInt(fileinfo.Size(), 36))
		ctx.Header(constant.LastModified, fileinfo.ModTime().Format(constant.TimeFormat))
	}
	protocol.GetHeader(ctx, content, res)
	return
}

func (posix *Posix) HeadObject(ctx *gin.Context) {
	res, bucket, key := protocol.Param(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	if err := posix.DirChecker(posix.getBucketPath(bucket)); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.NoSuchBucket)
		return
	}

	filename := filepath.Join(posix.getBucketPath(bucket), key)
	if err := posix.FileChecker(filename); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.NoSuchKey)
		return
	}

	fileinfo, err := os.Stat(filename)
	if err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.NoSuchKey)
		return
	} else {
		ctx.Header(constant.IsDir, fmt.Sprintf("%t", fileinfo.IsDir()))
		ctx.Header(constant.Mode, fmt.Sprintf("%d", fileinfo.Mode()))
		ctx.Header(constant.ContentLength, strconv.FormatInt(fileinfo.Size(), 10))
		ctx.Header(constant.LastModified, fileinfo.ModTime().Format(constant.TimeFormat))
	}
	return
}

func (posix *Posix) DeleteObject(ctx *gin.Context) {
	res, bucket, key := protocol.Param(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	if err := posix.DirChecker(posix.getBucketPath(bucket)); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.NoSuchBucket)
		return
	}

	filename := filepath.Join(posix.getBucketPath(bucket), key)
	if err := posix.FileChecker(filename); err != nil {
		res.Error(errors.NoSuchKey)
		return
	}

	err := os.Remove(filename)
	if err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.NoSuchKey)
		return
	}

	ctx.Status(http.StatusOK)
	return
}
