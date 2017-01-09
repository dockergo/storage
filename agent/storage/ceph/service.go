package ceph

import (
	"net/http"

	"flyaway.com/gin"
	"github.com/flyaways/storage/agent/util/log"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

func (c *ceph) Service(ctx *gin.Context) {
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

	for _, bucket := range resp.Buckets {
		c.XML(http.StatusOK, gin.H{"Bucket": bucket.Name})
	}
	ctx.Status(http.StatusOK)
}
