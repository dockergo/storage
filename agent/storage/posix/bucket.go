package posix

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/flyaways/storage/agent/errors"
	"github.com/flyaways/storage/agent/protocol"
	"github.com/flyaways/storage/agent/util/log"
)

func walkDir(dirPth string) (files []os.FileInfo, filenames []string, err error) {
	files = make([]os.FileInfo, 0, 30)
	filenames = make([]string, 0, 30)
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		fi.Name()
		files = append(files, fi)
		filenames = append(filenames, filename)
		return nil
	})

	return files, filenames, err
}

func (posix *Posix) GetBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}

	if err := posix.DirChecker(posix.getBucketPath(bucket)); err != nil {
		res.Error(errors.NoSuchBucket)
		ctx.Status(http.StatusNotFound)
		return
	}

	keys, filenames, err := walkDir(posix.getBucketPath(bucket))
	if err != nil {
		res.Error(errors.NoSuchBucket)
	}
	lens := len(posix.Config.Storage.Posix.Addr) + len(bucket) + 2
	for index, key := range keys {
		ctx.JSON(http.StatusOK, gin.H{"Key": filenames[index][lens:len(filenames[index])],
			"LastModified": key.ModTime(),
			"Size":         key.Size()})
	}
	ctx.Status(http.StatusOK)
}

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

	if err := posix.DirChecker(bucketPath); err != nil {
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

	if err := posix.DirChecker(posix.getBucketPath(bucket)); err != nil {
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

	if err := posix.DirChecker(bucketPath); err != nil {
		res.Error(errors.NoSuchBucket)
		return
	}

	f, err := os.Open(bucketPath)
	if err != nil {
		log.Warn("[%s:%s]", posix.Name, err.Error())
		res.Error(errors.NoSuchKey)
		return
	}

	list, err := f.Readdir(1)
	f.Close()

	if err != io.EOF {
		log.Warn("[%s:err != io.EOF]", posix.Name)
		res.Error(errors.BucketNotEmpty)
		return
	}
	if len(list) > 0 {
		log.Warn("[%s:len(list) > 0]", posix.Name)
		res.Error(errors.BucketNotEmpty)
		return
	}

	err = os.Remove(bucketPath)
	if err != nil {
		log.Warn("[%s:%s]", posix.Name, err.Error())
		res.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
