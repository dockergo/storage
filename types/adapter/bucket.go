package adapter

import (
	"net/http"

	"github.com/flyaways/storage/util"
	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
)

func (s *StorageAdapter) GetBucket(ctx *gin.Context) {
	log.Warn("[%s GetBucket NotImplemented]", s.Name)
	util.Details()
}

func (s *StorageAdapter) PutBucket(ctx *gin.Context) {
	log.Warn("[%s PutBucket NotImplemented]", s.Name)
	util.Details()
}

func (s *StorageAdapter) HeadBucket(ctx *gin.Context) {
	log.Warn("[%s HeadBucket NotImplemented]", s.Name)
	util.Details()
}

func (s *StorageAdapter) DeleteBucket(ctx *gin.Context) {
	log.Warn("[%s DeleteBucket NotImplemented]", s.Name)
	util.Details()
}

func (s *StorageAdapter) OptionsBucket(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Supported": "GET,PUT,HEAD,DELETE,OPTIONS"})
	ctx.Status(http.StatusOK)
}
