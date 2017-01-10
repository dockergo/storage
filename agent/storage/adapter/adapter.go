package adapter

import (
	"runtime"

	"github.com/gin-gonic/gin"

	"github.com/flyaways/storage/agent/util/log"
	"github.com/flyaways/tracker"
)

type Services interface {
	Service(ctx *gin.Context)
}

type Objecter interface {
	PutObject(ctx *gin.Context)
	PostObject(ctx *gin.Context)
	GetObject(ctx *gin.Context)
	HeadObject(ctx *gin.Context)
	DeleteObject(ctx *gin.Context)
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

func Details() {
	for skip := 0; true; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if ok {
			log.Warn(tracker.Blue("[%d\t%s\t%d\t%s]"), pc, file, line, runtime.FuncForPC(pc).Name())
		} else {
			break
		}
	}
}
