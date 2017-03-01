//Website: https://github.com/flyaways
//Modifier: Flyaway
//Date 27/09/2016 13:22 Beijing
//
package swift

import (
	"github.com/flyaways/storage/protocol"
	"github.com/gin-gonic/gin"
)

func (swt *Swift) GetBucket(ctx *gin.Context) {
	swt.bucket("GET", ctx)
}

func (swt *Swift) PostBucket(ctx *gin.Context) {
	swt.bucket("POST", ctx)
}

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
	res, bucket, _ := protocol.Param(ctx)
	if len(bucket) == 0 {
		return
	}
	url := buildBucketUrl(swt.config.Storage.Swift.Addr, swt.authAccount, bucket)
	swt.bucketrequest(nil, method, url, res, ctx)
}
