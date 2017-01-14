package oss

import (
	"net/http"

	"github.com/flyaways/storage/constant"
	"github.com/flyaways/storage/errors"
	"github.com/flyaways/storage/protocol"
	"github.com/flyaways/storage/util/log"
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

	meta, err := ossc.client.GetBucketInfo(bucket)
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.NoSuchBucket)
	}

	ctx.Header("Local", meta.XMLName.Local)
	ctx.Header("Space", meta.XMLName.Space)
	ctx.Header("Name", meta.BucketInfo.Name)
	ctx.Header("Location", meta.BucketInfo.Location)
	ctx.Header("CreationDate", meta.BucketInfo.CreationDate.Format(constant.TimeFormat))
	ctx.Header("ExtranetEndpoint", meta.BucketInfo.ExtranetEndpoint)
	ctx.Header("IntranetEndpoint", meta.BucketInfo.IntranetEndpoint)
	ctx.Header("ACL", meta.BucketInfo.ACL)
	ctx.Header("ID", meta.BucketInfo.Owner.ID)
	ctx.Header("DisplayName", meta.BucketInfo.Owner.DisplayName)

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
