package middleware

import (
	"time"

	"github.com/flyaways/storage/agent/constant"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqID := uuid.NewV4().String()
		ctx.Set("requestId", reqID)
		ctx.Header(constant.Server, "Nginx")
		ctx.Header(constant.Date, time.Now().UTC().Format(constant.TimeFormat))
		ctx.Header(constant.RequestId, reqID)
	}

}
