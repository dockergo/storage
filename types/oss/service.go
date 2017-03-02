package oss

import (
	"net/http"

	"github.com/flyaways/storage/errors"
	"github.com/flyaways/storage/result"
	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
)

func (ossc *OSS) ListBuckets(ctx *gin.Context) {
	res := result.NewResult(ctx)
	lsRes, err := ossc.client.ListBuckets()
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
		res.Error(errors.NoSuchBucket)
		ctx.Status(http.StatusNotFound)
	}

	for _, bucket := range lsRes.Buckets {
		ctx.JSON(http.StatusOK, gin.H{"Bucket": bucket.Name})
	}

}
