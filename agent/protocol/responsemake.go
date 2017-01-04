package protocol

import (
	"github.com/gin-gonic/gin"
	"github.com/flyaways/storage/agent/constant"
)

func responseMake(ctx *gin.Context) {
	if v := ctx.Query("response-content-type"); v != "" {
		ctx.Header(constant.ContentType, v)
	}
	if v := ctx.Query("response-content-language"); v != "" {
		ctx.Header(constant.ContentLanguage, v)
	}
	if v := ctx.Query("response-expires"); v != "" {
		ctx.Header(constant.Expires, v)
	}
	if v := ctx.Query("response-cache-control"); v != "" {
		ctx.Header(constant.CacheControl, v)
	}
	if v := ctx.Query("response-content-disposition"); v != "" {
		ctx.Header(constant.ContentDisposition, v)
	}
	if v := ctx.Query("response-content-encoding"); v != "" {
		ctx.Header(constant.ContentEncoding, v)
	}
}
