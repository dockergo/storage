package protocol

import (
	"fmt"
	"strconv"

	"github.com/flyaways/storage/constant"
	errs "github.com/flyaways/storage/errors"
	"github.com/flyaways/storage/result"
	"github.com/flyaways/storage/util"
	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
)

//符合 DNS 标准的存储桶名称规则如下：
//存储桶名称的长度必须为至少 3 个字符，且不能超过 63 个字符。
//存储桶名称必须是一系列的一个或多个标签。相邻标签通过单个句点 (.) 分隔。存储桶名称可以包含小写字母、数字和连字符。每个标签都必须以小写字母或数字开头和结尾。
//存储桶名称不得采用 IP 地址格式（例如，192.168.5.4）。
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

func ComputeKey(ctx *gin.Context, key string, data []byte) string {
	finalkey := util.GetSha1Hex(data)

	policy, exits := ctx.Get(constant.Policy)
	if exits {
		base, err := strconv.ParseInt(fmt.Sprintf("%d", policy), 10, 64)
		if err == nil {
			if (base%constant.Key - base%constant.Auth) == 4 {
				finalkey = key
			}
		}
	}
	return finalkey
}
