package kdfs

import (
	"github.com/flyaways/storage/agent/protocol"
	"github.com/gin-gonic/gin"
)

func (kfs *Kdfs) GetBucket(ctx *gin.Context) {
	kfs.bucket("GET", ctx)
}

func (kfs *Kdfs) PutBucket(ctx *gin.Context) {
	kfs.bucket("PUT", ctx)
}

func (kfs *Kdfs) HeadBucket(ctx *gin.Context) {
	kfs.bucket("HEAD", ctx)
}

func (kfs *Kdfs) DeleteBucket(ctx *gin.Context) {
	kfs.bucket("DELETE", ctx)
}

func (kfs *Kdfs) bucket(method string, ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}
	url := buildBucketUrl(kfs.config.Storage.Kdfs.Addr, kfs.config.Storage.Kdfs.Account, bucket)
	kfs.request(nil, method, url, res, ctx)
}
