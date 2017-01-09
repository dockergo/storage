package middleware

import (
	"strconv"
	"strings"

	"github.com/flyaways/storage/agent/constant"
	"github.com/gin-gonic/gin"
)

func Policy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bucket := ctx.Param("bucket")
		if len(bucket) < 3 {
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
