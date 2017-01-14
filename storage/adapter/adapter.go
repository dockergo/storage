package adapter

import "github.com/gin-gonic/gin"

type Services interface {
	ListBuckets(ctx *gin.Context)
}

type Objecter interface {
	PutObject(ctx *gin.Context)
	PostObject(ctx *gin.Context)
	GetObject(ctx *gin.Context)
	HeadObject(ctx *gin.Context)
	DeleteObject(ctx *gin.Context)
	MoveObject(ctx *gin.Context)
	CopyObject(ctx *gin.Context)
	OptionsObject(ctx *gin.Context)
}

type Bucketer interface {
	GetBucket(ctx *gin.Context)
	PutBucket(ctx *gin.Context)
	HeadBucket(ctx *gin.Context)
	DeleteBucket(ctx *gin.Context)
	OptionsBucket(ctx *gin.Context)
}

type Storager interface {
	Services
	Objecter
	Bucketer
}

type StorageAdapter struct {
	Name string
}
