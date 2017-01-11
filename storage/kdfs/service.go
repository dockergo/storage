package kdfs

import (
	"github.com/flyaways/storage/protocol"
	"github.com/gin-gonic/gin"
)

func (kfs *Kdfs) Service(ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	url := buildBucketUrl(kfs.config.Storage.Kdfs.Addr, kfs.config.Storage.Kdfs.Account, bucket)
	kfs.request(nil, "GET", url, res, ctx)
}
