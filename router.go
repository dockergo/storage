package storage

import (
	"github.com/flyaways/storage/app"
	"github.com/flyaways/storage/middleware"
	"github.com/gin-gonic/gin"
)

func regRouters(app *app.App, router *gin.Engine) {
	router.Use(middleware.RequestId())
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(middleware.Cors())
	router.Use(middleware.Sessions())
	router.Use(middleware.Csrf())

	greenGroup(app, router)
	bucketGroup(app, router)
	objectGroup(app, router)
	serviceGroup(app, router)
}

func greenGroup(app *app.App, router *gin.Engine) {
	green := router.Group("/:bucket")

	green.HEAD("", app.HeadBucket)
	green.HEAD("/*key", app.HeadObject)
	green.OPTIONS("/:bucket", app.OptionsBucket)
	green.OPTIONS("/:bucket/*key", app.OptionsObject)
}

func bucketGroup(app *app.App, router *gin.Engine) {
	bucket := router.Group("/:bucket")
	bucket.Use(middleware.Authority(app.Config.Credential))

	bucket.GET("", app.GetBucket)
	bucket.PUT("", app.PutBucket)
	bucket.DELETE("", app.DeleteBucket)

}

func objectGroup(app *app.App, router *gin.Engine) {
	object := router.Group("/:bucket")
	object.Use(middleware.Authority(app.Config.Credential))

	object.PUT("/*key", app.PutObject)
	object.POST("", app.PostObject)
	object.GET("/*key", app.GetObject)
	object.DELETE("/*key", app.DeleteObject)

}

func serviceGroup(app *app.App, router *gin.Engine) {
	service := router.Group("/")
	service.Use(middleware.Authority(app.Config.Credential))
	service.GET("", app.ListBuckets)

}
