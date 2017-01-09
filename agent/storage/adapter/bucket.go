package adapter

import (
	"net/http"

	"github.com/flyaways/storage/agent/util/log"
	"github.com/gin-gonic/gin"
)

func (s *StorageAdapter) GetBucket(ctx *gin.Context) {
	log.Warn("[%s GetBucket NotImplemented]", s.Name)
	Details()
}

func (s *StorageAdapter) PutBucket(ctx *gin.Context) {
	log.Warn("[%s PutBucket NotImplemented]", s.Name)
	Details()
}

func (s *StorageAdapter) HeadBucket(ctx *gin.Context) {
	log.Warn("[%s HeadBucket NotImplemented]", s.Name)
	Details()
}

func (s *StorageAdapter) DeleteBucket(ctx *gin.Context) {
	log.Warn("[%s DeleteBucket NotImplemented]", s.Name)
	Details()
}

func (s *StorageAdapter) OptionsBucket(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
