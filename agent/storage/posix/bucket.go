package posix

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/flyaways/storage/agent/errors"
	"github.com/flyaways/storage/agent/protocol"
	"github.com/flyaways/storage/agent/util/log"
)

func (posix *Posix) PutBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}

	err := os.MkdirAll(posix.getBucketPath(bucket), os.ModePerm)
	if err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.InvalidArgument)
		return
	}

	bucketPath := posix.getBucketPath(bucket)
	if err := posix.Checker.DirChecker(bucketPath); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.InvalidArgument)
		return
	}

	ctx.Status(http.StatusOK)
}

func (posix *Posix) HeadBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}

	if err := posix.Checker.DirChecker(posix.getBucketPath(bucket)); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.NoSuchBucket)
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusOK)
}

func (posix *Posix) DeleteBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}

	bucketPath := posix.getBucketPath(bucket)
	if err := posix.Checker.DirChecker(bucketPath); err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.NoSuchBucket)
		return
	}
	f, err := os.Open(bucketPath)
	if err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(err)
		return
	}

	list, err := f.Readdir(1)
	f.Close()

	if err != io.EOF {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.BucketNotEmpty)
		return
	}
	if len(list) > 0 {
		log.Error("[%s:len(list) > 0]", posix.Name)
		res.Error(errors.BucketNotEmpty)
		return
	}

	err = os.Remove(bucketPath)
	if err != nil {
		log.Error("[%s:%s]", posix.Name, err.Error())
		res.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
