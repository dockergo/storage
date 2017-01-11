package protocol

import (
	"net/http"
	"time"

	"github.com/flyaways/storage/constant"
	"github.com/gin-gonic/gin"
)

func CheckModSince(ctx *gin.Context, modtime time.Time) (bool, error) {
	if modtime.IsZero() {
		return false, nil
	}
	if v := ctx.Request.Header.Get(constant.IfModifiedSince); v != "" {
		modifiedSince, err := time.Parse(constant.TimeFormat, v)
		if err != nil {
			return true, err
		}

		if modtime.Before(modifiedSince) {
			ctx.Status(http.StatusNotModified)
			return true, nil
		}
	}

	if v := ctx.Request.Header.Get(constant.IfUnmodifiedSince); v != "" {
		umModifiedSince, err := time.Parse(constant.TimeFormat, v)
		if err != nil {
			return true, err
		}
		if modtime.After(umModifiedSince) {
			ctx.Status(http.StatusPreconditionFailed)
			return true, nil
		}
	}

	return false, nil
}
