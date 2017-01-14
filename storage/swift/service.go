package swift

import (
	"github.com/flyaways/storage/protocol"
	"github.com/gin-gonic/gin"
)

func (swt *Swift) ListBuckets(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	url := buildBucketUrl(swt.config.Storage.Swift.Addr, swt.authAccount, bucket)
	swt.request(nil, "GET", url, res, ctx)
}
