package oss

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flyaways/storage/constant"
	"github.com/flyaways/storage/errors"
	"github.com/flyaways/storage/protocol"
	"github.com/flyaways/storage/util"
	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
)

func (ossc *OSS) PutObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	data, key, err := protocol.PutHeadchecker(ctx, res, bucket, key)
	if err != nil {
		return
	}

	bucketobj, err := ossc.client.Bucket(bucket)
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.NoSuchBucket)
	}

	err = bucketobj.PutObject(key, bytes.NewReader(data))
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.UnSupportError)
	}

	ctx.Header(constant.ETag, util.GetETagValue(data))
	ctx.Header(constant.NewFileName, key)
	ctx.Status(http.StatusOK)
	ctx.JSON(http.StatusOK, gin.H{constant.NewFileName: key})
}

func (ossc *OSS) PostObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParamPost(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	data, key, err := protocol.PostHeadchecker(ctx, res, bucket, key)
	if err != nil {
		return
	}

	bucketobj, err := ossc.client.Bucket(bucket)
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.NoSuchBucket)
	}

	err = bucketobj.PutObject(key, bytes.NewReader(data))
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.UnSupportError)
	}

	ctx.Header(constant.ETag, util.GetETagValue(data))
	ctx.Header(constant.NewFileName, key)
	ctx.Status(http.StatusOK)
	ctx.JSON(http.StatusOK, gin.H{constant.NewFileName: key})
}

func (ossc *OSS) GetObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	bucketobj, err := ossc.client.Bucket(bucket)
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.NoSuchBucket)
	}

	body, err := bucketobj.GetObject(key)
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.NoSuchBucket)
	}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.UnSupportError)
	}
	body.Close()
	fmt.Println("data:", string(data))
}

func (ossc *OSS) DeleteObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}
	bucketobj, err := ossc.client.Bucket(bucket)
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.NoSuchBucket)
	}

	err = bucketobj.DeleteObject(key)
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.UnSupportError)
	}
}

func (ossc *OSS) HeadObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	bucketobj, err := ossc.client.Bucket(bucket)
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.NoSuchBucket)
	}

	header, err := bucketobj.GetObjectMeta(key)
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.UnSupportError)
	}

	objEtag := header.Get(constant.ETag)
	if protocol.CheckETag(ctx, objEtag) {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.UnSupportError)
		return
	}
	objLastModified := header.Get(constant.LastModified)
	protocol.HeadChecker(ctx, res, objEtag, objLastModified)

}
