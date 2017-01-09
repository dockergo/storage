package adapter

import (
	"flyaway.com/gin"
	"github.com/flyaways/storage/agent/util/log"
)

func (s *StorageAdapter) Service(ctx *gin.Context) {
	log.Warn("[%s Service NotImplemented]", s.Name)
	Details()
}
