package storage

import (
	"net/http"

	"github.com/flyaways/storage/app"
	"github.com/flyaways/storage/middleware"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func RegisterURLs(app *app.App, router *gin.Engine) {
	router.Use(middleware.RequestId())
	router.Use(middleware.Cors())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//noAuth
	noAuth := router.Group("/:bucket")
	{
		noAuth.HEAD("", app.HeadBucket)
		noAuth.HEAD("/*key", app.HeadObject)
		router.OPTIONS("/:bucket", app.OptionsBucket)
		router.OPTIONS("/:bucket/*key", app.OptionsObject)
	}

	store := sessions.NewCookieStore([]byte("secret"))

	//authRequired
	authRequired := router.Group("/:bucket")

	authRequired.Use(middleware.AuthRequired(app.Config.Credential))
	authRequired.Use(sessions.Sessions("storagesession", store))
	authRequired.Use(csrf.Middleware(csrf.Options{
		Secret: "storagesecret",
		ErrorFunc: func(c *gin.Context) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "CSRF token mismatch",
			})
			c.Abort()
		},
	}))

	{
		authRequired.GET("", app.GetBucket)
		authRequired.PUT("", app.PutBucket)
		authRequired.DELETE("", app.DeleteBucket)

		authRequired.PUT("/*key", app.PutObject)
		authRequired.POST("", app.PostObject)
		authRequired.GET("/*key", app.GetObject)
		authRequired.DELETE("/*key", app.DeleteObject)
	}

	//service
	service := router.Group("/")

	service.Use(middleware.AuthRequired(app.Config.Credential))
	service.Use(sessions.Sessions("storagesession", store))
	service.Use(csrf.Middleware(csrf.Options{
		Secret: "storagesecret",
		ErrorFunc: func(c *gin.Context) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "CSRF token mismatch",
			})
			c.Abort()
		},
	}))

	{
		service.GET("", app.ListBuckets)
	}

}
