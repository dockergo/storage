package kdfs

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

func (kfs *Kdfs) PutObject(ctx *gin.Context) {
	kfs.object("PUT", ctx)
}

func (kfs *Kdfs) PostObject(ctx *gin.Context) {
	kfs.object("PUT", ctx)
}

func (kfs *Kdfs) HeadObject(ctx *gin.Context) {
	kfs.object("HEAD", ctx)
}

func (kfs *Kdfs) GetObject(ctx *gin.Context) {
	kfs.object("GET", ctx)
}

func (kfs *Kdfs) DeleteObject(ctx *gin.Context) {
	kfs.object("DELETE", ctx)
}

func (kfs *Kdfs) object(method string, ctx *gin.Context) {
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

	url := buildUrl(kfs.config.Storage.Kdfs.Addr, kfs.config.Storage.Kdfs.Account, bucket, key)
	kfs.request(bytes.NewReader(data), method, url, res, ctx)

	ctx.Header(constant.ETag, util.GetETagValue(data))
	ctx.Header(constant.NewFileName, key)
	ctx.JSON(http.StatusOK, gin.H{constant.NewFileName: key})
}
