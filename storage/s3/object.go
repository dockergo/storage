package s3

import (
	"net/http"

	"github.com/flyaways/storage/constant"
	"github.com/flyaways/storage/protocol"
	"github.com/flyaways/storage/result"
	"github.com/flyaways/storage/util"
	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/goamz/s3"
)

func (c *S3c) PutObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	data, finalkey, err := protocol.PutHeader(ctx, res, bucket, key)
	if err != nil {
		return
	}

	c.upload(ctx, res, data, finalkey, bucket)
}

func (c *S3c) PostObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParamPost(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	data, finalkey, err := protocol.PostHeader(ctx, res, bucket, key)
	if err != nil {
		return
	}

	c.upload(ctx, res, data, finalkey, bucket)
}

func (c *S3c) upload(ctx *gin.Context, res *result.Result, data []byte, finalkey, bucket string) {
	v := ctx.Request.Header.Get(constant.ContentType)
	err := c.client.Bucket(bucket).Put(finalkey, data, v, s3.PublicReadWrite)
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}

	ctx.Header(constant.ETag, util.GetETagValue(data))
	ctx.Header(constant.NewFileName, finalkey)
	ctx.Status(http.StatusOK)
	ctx.JSON(http.StatusOK, gin.H{constant.NewFileName: finalkey})
}

func (c *S3c) GetObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	body, err := c.client.Bucket(bucket).Get(key)
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}

	protocol.GetHeader(ctx, body, res)
}

func (c *S3c) HeadObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	httpRep, err := c.client.Bucket(bucket).Head(key)
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}

	for key, value := range httpRep.Header {
		for _, values := range value {
			ctx.Header(key, values)
		}
	}

	ctx.Status(httpRep.StatusCode)
}

func (c *S3c) DeleteObject(ctx *gin.Context) {
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
