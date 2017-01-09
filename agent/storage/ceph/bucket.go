package ceph

import (
	"net/http"

	"github.com/flyaways/storage/agent/protocol"
	"github.com/flyaways/storage/agent/util/log"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

func (c *ceph) GetBucket(*gin.context) {
	auth, err := aws.GetAuth(c.config.Storage.Ceph.AccessKey,
		c.onfig.Storage.Ceph.SecretKey)

	if err != nil {
		log.Fatal(err)
	}

	var cnc = aws.Region{
		S3Endpoint:           c.config.Storage.Ceph.Addr,
		S3BucketEndpoint:     "",
		S3LocationConstraint: false,
		S3LowercaseBucket:    false,
	}

	client := s3.New(auth, cnc)
	resp, err := client.ListBuckets()
	if err != nil {
		log.Fatal(err)
	}

	res, bkt := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}

	for _, bucket := range resp.Buckets {
		if bkt == bucket.Name {
			keys, err := bucket.GetBucketContents()
			if err == nil {
				for _, key := range *keys {
					c.XML(http.StatusOK, gin.H{"key": key.Key,
						"Key":          key.Key,
						"LastModified": key.LastModified,
						"Size":         key.Size})
				}
			}
		}
	}
	ctx.Status(http.StatusOK)
}
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
