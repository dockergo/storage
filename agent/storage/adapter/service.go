package adapter

import (
	"github.com/flyaways/storage/agent/util/log"
	"github.com/gin-gonic/gin"
)

func (s *StorageAdapter) Service(ctx *gin.Context) {
	log.Warn("[%s Service NotImplemented]", s.Name)
	Details()
}
