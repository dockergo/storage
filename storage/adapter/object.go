package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/flyaways/storage/util"
	"github.com/flyaways/storage/util/log"
)

func (s *StorageAdapter) PutObject(ctx *gin.Context) {
	log.Warn("[%s PutObject NotImplemented]", s.Name)
	util.Details()
}

func (s *StorageAdapter) PostObject(ctx *gin.Context) {
	log.Warn("[%s PostObject NotImplemented]", s.Name)
	util.Details()
}

func (s *StorageAdapter) GetObject(ctx *gin.Context) {
	log.Warn("[%s GetObject NotImplemented]", s.Name)
	util.Details()
}

func (s *StorageAdapter) HeadObject(ctx *gin.Context) {
	log.Warn("[%s HeadObject NotImplemented]", s.Name)
	util.Details()
}

func (s *StorageAdapter) DeleteObject(ctx *gin.Context) {
	log.Warn("[%s DeleteObject NotImplemented]", s.Name)
	util.Details()
}

func (s *StorageAdapter) MoveObject(ctx *gin.Context) {
	log.Warn("[%s MoveObject NotImplemented]", s.Name)
	util.Details()
}

func (s *StorageAdapter) CopyObject(ctx *gin.Context) {
	log.Warn("[%s CopyObject NotImplemented]", s.Name)
	util.Details()
}

func (s *StorageAdapter) OptionsObject(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Supported": "GET,PUT,POST,HEAD,DELETE,OPTIONS"})
	ctx.Status(http.StatusOK)
}
