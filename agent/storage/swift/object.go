//Website: https://github.com/flyaways
//Modifier: Flyaway
//Date 27/09/2016 13:22 Beijing
//
package swift

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/flyaways/storage/agent/constant"
	"github.com/flyaways/storage/agent/errors"
	"github.com/flyaways/storage/agent/protocol"
	"github.com/flyaways/storage/agent/util/log"
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

	url := buildUrl(swt.config.Storage.Swift.Addr, swt.authAccount, bucket, key)
	swt.request(data, method, url, res, ctx)
	ctx.JSON(http.StatusOK, gin.H{constant.NewFileName: key})
}
