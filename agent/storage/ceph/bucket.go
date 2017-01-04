package ceph

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goamz/goamz/s3"
	"github.com/flyaways/storage/agent/protocol"
	"github.com/flyaways/storage/agent/util/log"
)

func (c *Ceph) PutBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}

	if err := c.client.Bucket(bucket).PutBucket(s3.PublicRead); err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *Ceph) HeadBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}

	resp, err := c.client.Bucket(bucket).Head("/", ctx.Request.Header)
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}

	ctx.Status(resp.StatusCode)
}

func (c *Ceph) DeleteBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}

	if err := c.client.Bucket(bucket).DelBucket(); err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
