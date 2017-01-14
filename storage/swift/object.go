//Website: https://github.com/flyaways
//Modifier: Flyaway
//Date 27/09/2016 13:22 Beijing
//
package swift

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/flyaways/storage/constant"
	"github.com/flyaways/storage/errors"
	"github.com/flyaways/storage/protocol"
	"github.com/flyaways/storage/util"
	"github.com/flyaways/storage/util/log"
)

func (swt *Swift) PutObject(ctx *gin.Context) {
	swt.object("PUT", ctx)
}

func (swt *Swift) PostObject(ctx *gin.Context) {
	swt.object("PUT", ctx)
}

func (swt *Swift) HeadObject(ctx *gin.Context) {
	swt.object("HEAD", ctx)
}

func (swt *Swift) GetObject(ctx *gin.Context) {
	swt.object("GET", ctx)
}

func (swt *Swift) DeleteObject(ctx *gin.Context) {
	swt.object("DELETE", ctx)
}

func (swt *Swift) object(method string, ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	var data []byte
	var err error
	if method == "PUT" {
		data, key, err = protocol.PutHeader(ctx, res, bucket, key)
	} else if method == "POST" {
		data, key, err = protocol.PostHeader(ctx, res, bucket, key)
		if err != nil {
			log.Error("[object.POST read multipart error:%s]", err.Error())
			res.Error(errors.InternalError)
			return
		}
	}

	url := buildUrl(swt.config.Storage.Swift.Addr, swt.authAccount, bucket, key)
	swt.request(bytes.NewReader(data), method, url, res, ctx)

	ctx.Header(constant.ETag, util.GetETagValue(data))
	ctx.Header(constant.NewFileName, key)
	ctx.JSON(http.StatusOK, gin.H{constant.NewFileName: key})
}
