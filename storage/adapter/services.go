package adapter

import (
	"github.com/flyaways/storage/util"
	"github.com/flyaways/storage/util/log"
	"github.com/gin-gonic/gin"
)

func (s *StorageAdapter) ListBuckets(ctx *gin.Context) {
	log.Warn("[%s ListBuckets NotImplemented]", s.Name)
	util.Details()
}
