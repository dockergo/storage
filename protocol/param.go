package protocol

import (
	"net/http"

	errs "github.com/flyaways/storage/errors"
	"github.com/flyaways/storage/result"
	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
)

func Param(ctx *gin.Context) (*result.Result, string, string) {
	res := result.NewResult(ctx)
	bucket := ctx.Param("bucket")
	var key string
	var exists bool

	if ctx.Request.Method == "POST" {
		key, exists = ctx.GetPostForm("key")
		if !exists {
			log.Error("[%s,%s]", bucket, key)
			res.Error(errs.InvalidArgument)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return res, "", ""
		}
	} else {
		key = ctx.Param("key")
	}

	if key != "" {
		key = key[1:len(key)]
	}

	if len(bucket) < 3 || len(bucket) > 63 {
		log.Error("[%s/%s]", bucket, key)
		res.Error(errs.InvalidArgument)
		return res, "", ""
	}

	return res, bucket, key
}
