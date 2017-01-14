package swift

import (
	"github.com/flyaways/storage/protocol"
	"github.com/gin-gonic/gin"
)

func (swt *Swift) ListBuckets(ctx *gin.Context) {
	res, _, _ := protocol.Param(ctx)
	url := buildBucketUrl(swt.config.Storage.Swift.Addr, swt.authAccount, "")
	swt.request(nil, "GET", url, res, ctx)
}
