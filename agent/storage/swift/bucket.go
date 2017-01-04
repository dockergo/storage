//Website: https://github.com/flyaways
//Modifier: Flyaway
//Date 27/09/2016 13:22 Beijing
//
package swift

import (
	"github.com/gin-gonic/gin"
	"github.com/flyaways/storage/agent/protocol"
)

func (swt *Swift) PutBucket(ctx *gin.Context) {
	swt.bucket("PUT", ctx)
}

func (swt *Swift) HeadBucket(ctx *gin.Context) {
	swt.bucket("HEAD", ctx)
}

func (swt *Swift) DeleteBucket(ctx *gin.Context) {
	swt.bucket("DELETE", ctx)
}

func (swt *Swift) bucket(method string, ctx *gin.Context) {
	res, bucket := protocol.GetParamBucket(ctx)
	if len(bucket) == 0 {
		return
	}
	url := buildBucketUrl(swt.config.Storage.Swift.Addr, swt.authAccount, bucket)
	swt.request(nil, method, url, res, ctx)
}
