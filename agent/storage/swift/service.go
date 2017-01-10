package swift

import (
	"github.com/flyaways/storage/agent/protocol"
	"github.com/gin-gonic/gin"
)

func (swt *Swift) Service(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	url := buildBucketUrl(swt.config.Storage.Swift.Addr, swt.authAccount, bucket)
	swt.request(nil, "GET", url, res, ctx)
}
