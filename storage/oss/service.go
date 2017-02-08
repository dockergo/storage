package oss

import (
	"net/http"

	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
)

func (ossc *OSS) ListBuckets(ctx *gin.Context) {
	lsRes, err := ossc.client.ListBuckets()
	if err != nil {
		log.Error("[%s:%s]", ossc.Name, err.Error())
	}

	for _, bucket := range lsRes.Buckets {
		ctx.JSON(http.StatusOK, gin.H{"Bucket": bucket.Name})
	}

}