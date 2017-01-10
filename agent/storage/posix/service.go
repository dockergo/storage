package posix

import (
	"io/ioutil"
	"net/http"

	"github.com/flyaways/storage/agent/errors"
	"github.com/flyaways/storage/agent/protocol"
	"github.com/flyaways/storage/agent/util/log"
	"github.com/gin-gonic/gin"
)

func walkDirs(path string) (buckets []string, err error) {
	buckets = make([]string, 0, 30)
	files, err := ioutil.ReadDir(path)
	for _, fi := range files {
		if fi.IsDir() {
			buckets = append(buckets, fi.Name())
		}
	}
	return buckets, err
}

func (posix *Posix) Service(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if err := posix.DirChecker(posix.getBucketPath(bucket)); err != nil {
		res.Error(errors.NoSuchBucket)
		ctx.Status(http.StatusNotFound)
		return
	}

	buckets, err := walkDirs(posix.Config.Storage.Posix.Addr)
	if err != nil {
		log.Error("[listbucket:%s]", err.Error())
		res.Error(errors.NoSuchBucket)
	}

	for _, bucket := range buckets {
		ctx.JSON(http.StatusOK, gin.H{"Bucket": bucket})
	}
	ctx.Status(http.StatusOK)
}
