package protocol

import (
	errs "github.com/flyaways/storage/errors"
	"github.com/flyaways/storage/result"
	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
)

func GetParam(ctx *gin.Context) (*result.Result, string, string) {
	res := result.NewResult(ctx)
	bucket := ctx.Param("bucket")
	key := ctx.Param("key")
	key = key[1:len(key)]
	if len(bucket) < 3 || len(bucket) > 63 || len(key) == 0 {
		log.Error("[%s/%s]", bucket, key)
		res.Error(errs.InvalidArgument)
		return res, "", ""
	}
	return res, bucket, key
}

func GetParamPost(ctx *gin.Context) (*result.Result, string, string) {
	res := result.NewResult(ctx)
	bucket := ctx.Param("bucket")
	key, exists := ctx.GetPostForm("key")
	key = key[1:len(key)]
	if len(bucket) < 3 || len(bucket) > 63 || !exists {
		log.Error("[%s,%s]", bucket, key)
		res.Error(errs.InvalidArgument)
		return res, "", ""
	}
	return res, bucket, key
}

func GetParamBucket(ctx *gin.Context) (*result.Result, string) {
	bucket := ctx.Param("bucket")
	res := result.NewResult(ctx)
	return res, bucket
}
