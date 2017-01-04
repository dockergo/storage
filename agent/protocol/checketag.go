package protocol

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/flyaways/storage/agent/constant"
)

func CheckETag(ctx *gin.Context, etag string) bool {
	if etag == "" {
		return false
	}
	if v := ctx.Request.Header.Get(constant.IfMatch); v != "" {
		if v != etag {
			ctx.Status(http.StatusPreconditionFailed)
			return true
		}
	}
	if v := ctx.Request.Header.Get(constant.IfNoneMatch); v != "" {
		if etag == v {
			ctx.Status(http.StatusNotModified)
			return true
		}
	}
	return false
}
