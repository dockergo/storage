package qiniu

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/flyaways/storage/constant"
	"github.com/flyaways/storage/errors"
	"github.com/flyaways/storage/protocol"
	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

func (c *Qiniu) PutObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	policy := &kodo.PutPolicy{
		Scope:   bucket + ":" + key,
		Expires: 3600,
	}

	token := c.client.MakeUptoken(policy)
	uploader := kodocli.NewUploader(0, nil)

	data, key, err := protocol.PutHeader(ctx, res, bucket, key)
	if err != nil {
		return
	}

	var ret PutRet
	err = uploader.Put(nil, &ret, token, key, bytes.NewReader(data), ctx.Request.ContentLength, nil)
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}
	ctx.Status(http.StatusOK)
}

func (c *Qiniu) PostObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParamPost(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	policy := &kodo.PutPolicy{
		Scope:   bucket + ":" + key,
		Expires: 3600,
	}

	token := c.client.MakeUptoken(policy)
	uploader := kodocli.NewUploader(0, nil)

	data, key, err := protocol.PutHeader(ctx, res, bucket, key)
	if err != nil {
		return
	}

	var ret PutRet
	err = uploader.Put(nil, &ret, token, key, bytes.NewReader(data), ctx.Request.ContentLength, nil)
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}
	ctx.Status(http.StatusOK)
}

func (c *Qiniu) HeadObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	entry, err := c.client.Bucket(bucket).Stat(nil, key)
	if err != nil {
		log.Error("[%s:%s]", entry, err.Error())
		res.Error(err)
		return
	}

	ctx.Header(constant.Hash, entry.Hash)
	ctx.Header(constant.EndUser, entry.EndUser)
	ctx.Header(constant.MimeType, entry.MimeType)
	ctx.Header(constant.ContentLength, strconv.FormatInt(entry.Fsize, 10))
	ctx.Header(constant.LastModified, time.Unix(entry.PutTime, 0).Format(constant.TimeFormat))
	ctx.Status(http.StatusOK)
}

func (c *Qiniu) DeleteObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	err := c.client.Bucket(bucket).Delete(nil, key)
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}
	ctx.Status(http.StatusOK)
}

func (c *Qiniu) GetObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	policy := kodo.GetPolicy{}
	downloadurl := c.client.MakePrivateUrl(kodo.MakeBaseUrl(c.config.Storage.Qiniu.Addr, key), &policy)

	file, err := os.Open(downloadurl)
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(errors.UnknownError)
		return
	}

	defer file.Close()
	fileName := url.QueryEscape(path.Base(downloadurl))
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("content-disposition", "attachment; filename=\""+fileName+"\"")

	_, err = io.Copy(ctx.Writer, file)
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(errors.UnknownError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *Qiniu) MoveObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	err := c.client.Bucket(bucket).Move(nil, key, movekey)
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}
	ctx.Status(http.StatusOK)
}

func (c *Qiniu) CopyObject(ctx *gin.Context) {
	res, bucket, key := protocol.GetParam(ctx)
	if len(bucket) == 0 || len(key) == 0 {
		return
	}

	err := c.client.Bucket(bucket).Copy(nil, key, copykey)
	if err != nil {
		log.Error("[%s:%s]", c.Name, err.Error())
		res.Error(err)
		return
	}
	ctx.Status(http.StatusOK)
}
