package s3

import (
	"net/http"

	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
)

func (c *S3c) ListBuckets(ctx *gin.Context) {
	resp, err := c.client.ListBuckets()
	if err != nil {
		log.Error("[listbucket:%s]", err.Error())
	}

	for _, bucket := range resp.Buckets {
		ctx.JSON(http.StatusOK, gin.H{"Bucket": bucket.Name})
	}
	ctx.Status(http.StatusOK)
}
