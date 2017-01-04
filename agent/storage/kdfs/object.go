package kdfs

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/flyaways/storage/agent/constant"
	"github.com/flyaways/storage/agent/errors"
	"github.com/flyaways/storage/agent/protocol"
	"github.com/flyaways/storage/agent/util/log"
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

	var data io.Reader
	var err error
	if method == "PUT" {
		data = ctx.Request.Body
	} else if method == "POST" {
		data, _, err = ctx.Request.FormFile("file")
		if err != nil {
			log.Error("[object.POST read multipart error:%s]", err.Error())
			res.Error(errors.InternalError)
			return
		}
	} else {
		data = nil
	}

	url := buildUrl(kfs.config.Storage.Kdfs.Addr, kfs.config.Storage.Kdfs.Account, bucket, key)
	kfs.request(data, method, url, res, ctx)
	ctx.JSON(http.StatusOK, gin.H{constant.NewFileName: key})
}
