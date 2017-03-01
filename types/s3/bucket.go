package s3

import (
	"net/http"

	"github.com/flyaways/storage/protocol"
	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/goamz/s3"
)

func (c *S3c) GetBucket(ctx *gin.Context) {
	resp, err := c.client.ListBuckets()
	if err != nil {
		log.Error("[listbucket:%s]", err.Error())
	}

	res, bkt, _ := protocol.Param(ctx)

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
				res.Error(err)
			}
		}
	}
	ctx.Status(http.StatusOK)
}

func (c *S3c) PutBucket(ctx *gin.Context) {
	res, bucket, _ := protocol.Param(ctx)
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

func (c *S3c) HeadBucket(ctx *gin.Context) {
	res, bucket, _ := protocol.Param(ctx)
	if len(bucket) == 0 {
		return
	}

	resp, err := c.client.Bucket(bucket).Head("/")
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}

	for key, value := range resp.Header {
		for _, values := range value {
			ctx.Header(key, values)
		}
	}

	ctx.Status(resp.StatusCode)
}

func (c *S3c) DeleteBucket(ctx *gin.Context) {
	res, bucket, _ := protocol.Param(ctx)
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
