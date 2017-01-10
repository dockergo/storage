package oss

import (
	"net/http"

	"github.com/flyaways/storage/agent/errors"
	"github.com/flyaways/storage/agent/protocol"
	"github.com/flyaways/storage/agent/util/log"
	"github.com/gin-gonic/gin"
)

func (ossc *OSS) PutBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}
	err := ossc.client.CreateBucket(bucket)
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.UnSupportError)
	}

}

func (ossc *OSS) PostBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}
	err := ossc.client.CreateBucket(bucket)
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.UnSupportError)
	}

}

func (ossc *OSS) DeleteBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}

	err := ossc.client.DeleteBucket(bucket)
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.BucketNotEmpty)
	}
}

func (ossc *OSS) HeadBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}

	status, err := ossc.client.IsBucketExist(bucket)
	if err != nil || !status {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.NoSuchBucket)
	}
}

func (ossc *OSS) GetBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}
	bucketobj, err := ossc.client.Bucket(bucket)
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.NoSuchBucket)
	}

	lsRes, err := bucketobj.ListObjects()
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.NoSuchBucket)
	}

	for _, object := range lsRes.Objects {
		ctx.JSON(http.StatusOK, gin.H{"Key": object.Key,
			"LastModified": object.LastModified,
			"Size":         object.Size})
	}
}
