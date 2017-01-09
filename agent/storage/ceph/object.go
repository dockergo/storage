package ceph

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/flyaways/storage/agent/constant"
	"github.com/flyaways/storage/agent/protocol"
	"github.com/flyaways/storage/agent/result"
	"github.com/flyaways/storage/agent/util"
	"github.com/flyaways/storage/agent/util/log"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/goamz/s3"
)

func (c *Ceph) PutObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	rawRequestdata, finalkey, err := protocol.PutHeadchecker(ctx, res, bucket, key)
	if err != nil {
		return
	}

	c.uploadObject(ctx, res, rawRequestdata, finalkey, bucket)
}

func (c *Ceph) PostObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParamPost(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	rawRequestdata, finalkey, err := protocol.PostHeadchecker(ctx, res, bucket, key)
	if err != nil {
		return
	}

	c.uploadObject(ctx, res, rawRequestdata, finalkey, bucket)
}

func (c *Ceph) uploadObject(ctx *gin.Context, res *result.Result, rawRequestdata []byte, finalkey, bucket string) {
	v := ctx.Request.Header.Get(constant.ContentType)
	err := c.client.Bucket(bucket).Put(finalkey, rawRequestdata, v, s3.PublicReadWrite, s3.Options{})
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}

	ctx.Header(constant.ETag, util.GetETagValue(rawRequestdata))
	ctx.Header(constant.NewFileName, finalkey)
	ctx.Status(http.StatusOK)
	ctx.JSON(http.StatusOK, gin.H{constant.NewFileName: finalkey})
}

func (c *Ceph) GetObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	rc, err := c.client.Bucket(bucket).GetReader(key)
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}
	content, _ := ioutil.ReadAll(rc)
	protocol.GetCkecker(ctx, content, res)
}

func (c *Ceph) HeadObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	resp, err := c.client.Bucket(bucket).Head(key, ctx.Request.Header)
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}
	//no check
	//objEtag := ctx.Request.Header.Get(constant.ETag)
	objEtag := resp.Header.Get(constant.ETag)
	if protocol.CheckETag(ctx, objEtag) {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(errors.New("invalid Etag"))
		return
	}
	objLastModified := resp.Header.Get(constant.LastModified)
	protocol.HeadChecker(ctx, res, objEtag, objLastModified)
	ctx.Status(resp.StatusCode)
}

func (c *Ceph) DeleteObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	if err := c.client.Bucket(bucket).Del(key); err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
