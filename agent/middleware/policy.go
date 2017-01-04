package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/flyaways/storage/agent/constant"
	"github.com/flyaways/storage/agent/util/log"
)

func Policy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bucket := ctx.Param("bucket")
		if len(bucket) < 3 {
			log.Warn("[len(%s) < 3]", bucket)
			return
		}

		if strings.HasPrefix(bucket, "bk") {
			policy := bucket[2:3]
			base, err := strconv.ParseInt(policy, 16, 64)
			if err == nil {
				ctx.Set(constant.Policy, base)
			}
		}
	}
}
