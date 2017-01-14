package storage

import (
	"github.com/flyaways/storage/app"
	"github.com/flyaways/storage/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterURLs(app *app.App, router *gin.Engine) {
	router.Use(middleware.RequestId())
	router.Use(middleware.Sessions())
	router.Use(middleware.Csrf())
	router.Use(middleware.Cors())
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	greenGroup(app, router)

	bucketGroup(app, router)

	objectGroup(app, router)

	serviceGroup(app, router)
}

func greenGroup(app *app.App, router *gin.Engine) {
	noAuth := router.Group("/:bucket")
	{
		noAuth.HEAD("", app.HeadBucket)
		noAuth.HEAD("/*key", app.HeadObject)
		router.OPTIONS("/:bucket", app.OptionsBucket)
		router.OPTIONS("/:bucket/*key", app.OptionsObject)
	}
}

func bucketGroup(app *app.App, router *gin.Engine) {
	required := router.Group("/:bucket")
	required.Use(middleware.AuthRequired(app.Config.Credential))

	required.GET("", app.GetBucket)
	required.PUT("", app.PutBucket)
	required.DELETE("", app.DeleteBucket)

}

func objectGroup(app *app.App, router *gin.Engine) {
	required := router.Group("/:bucket")
	required.Use(middleware.AuthRequired(app.Config.Credential))

	required.PUT("/*key", app.PutObject)
	required.POST("", app.PostObject)
	required.GET("/*key", app.GetObject)
	required.DELETE("/*key", app.DeleteObject)

}

func serviceGroup(app *app.App, router *gin.Engine) {
	service := router.Group("/")
	service.Use(middleware.AuthRequired(app.Config.Credential))

	{
		service.GET("", app.ListBuckets)
	}

}
