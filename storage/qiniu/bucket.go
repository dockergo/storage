package qiniu

import (
	"net/http"
	"strconv"
	"time"

	log "qiniupkg.com/x/log.v7"

	"github.com/flyaways/storage/constant"
	"github.com/flyaways/storage/errors"
	"github.com/flyaways/storage/protocol"
	"github.com/gin-gonic/gin"
)

func (c *Qiniu) GetBucket(ctx *gin.Context) {
	res, bucket, _ := protocol.Param(ctx)
	if len(bucket) == 0 {
		return
	}

	ListItem, _, _, err := c.client.Bucket(bucket).List(nil, "", "", "", 1000)
	if err != nil {
		log.Error("[%s]", err.Error())
		res.Error(errors.UnknownError)
	}

	for _, item := range ListItem {
		ctx.JSON(http.StatusOK, gin.H{"Key": item.Key,
			constant.LastModified: time.Unix(item.PutTime, 0).Format(constant.TimeFormat),
			constant.Size:         strconv.FormatInt(item.Fsize, 10),
			constant.Hash:         item.Hash,
			constant.EndUser:      item.EndUser,
			constant.MimeType:     item.MimeType})

	}
}

func (c *Qiniu) HeadBucket(ctx *gin.Context) {

}
