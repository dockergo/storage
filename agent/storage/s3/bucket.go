package s3

import (
	"net/http"

	"github.com/flyaways/storage/agent/protocol"
	"github.com/flyaways/storage/agent/util/log"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/goamz/s3"
)

func (c *s3c) GetBucket(ctx *gin.Context) {
	resp, err := c.client.ListBuckets()
	if err != nil {
		log.Error("[listbucket:%s]", err.Error())
	}

	_, bkt := protocol.GetParamBucket(ctx)
	if len(bkt) == 0 {
		return
	}

	for _, bucket := range resp.Buckets {
		if bkt == bucket.Name {
			keys, err := bucket.GetBucketContents()
			if err == nil {
				for _, key := range *keys {
					ctx.JSON(http.StatusOK, gin.H{"Key": key.Key,
						"LastModified": key.LastModified,
						"Size":         key.Size})
				}
			} else {
				log.Error("[%s]", err.Error())
			}
		}
	}
	ctx.Status(http.StatusOK)
}

func (c *s3c) PutBucket(ctx *gin.Context) {
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

func (c *s3c) HeadBucket(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}

	resp, err := c.client.Bucket(bucket).Head("/")
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}

	ctx.Status(resp.StatusCode)
}

func (c *s3c) DeleteBucket(ctx *gin.Context) {
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
