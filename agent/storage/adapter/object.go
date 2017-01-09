package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/flyaways/storage/agent/util/log"
)

func (s *StorageAdapter) PutObject(ctx *gin.Context) {
	log.Warn("[%s PutObject NotImplemented]", s.Name)
	Details()
}

func (s *StorageAdapter) GetObject(ctx *gin.Context) {
	log.Warn("[%s GetObject NotImplemented]", s.Name)
	Details()
}

func (s *StorageAdapter) HeadObject(ctx *gin.Context) {
	log.Warn("[%s HeadObject NotImplemented]", s.Name)
	Details()
}

func (s *StorageAdapter) DeleteObject(ctx *gin.Context) {
	log.Warn("[%s DeleteObject NotImplemented]", s.Name)
	Details()
}

func (s *StorageAdapter) OptionsObject(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Supported": "GET,PUT,POST,HEAD,DELETE,OPTIONS"})
	ctx.Status(http.StatusOK)
}
