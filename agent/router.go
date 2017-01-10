package agent

import (
	"net/http"

	"github.com/flyaways/storage/agent/app"
	"github.com/flyaways/storage/agent/middleware"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func RegisterURLs(app *app.App, router *gin.Engine) {
	router.Use(middleware.RequestId())
	router.Use(middleware.Cors())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.Policy())

	router.OPTIONS("/:bucket", app.OptionsBucket)
	router.OPTIONS("/:bucket/*key", app.OptionsObject)

	router.Use(middleware.AuthRequired(app.Config.Credential))

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("agentsession", store))
	router.Use(csrf.Middleware(csrf.Options{
		Secret: "agentsecret",
		ErrorFunc: func(c *gin.Context) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "CSRF token mismatch",
			})
			c.Abort()
		},
	}))

	bucket := router.Group("/:bucket")
	{
		bucket.GET("", app.GetBucket)
		bucket.PUT("", app.PutBucket)
		bucket.HEAD("", app.HeadBucket)
		bucket.DELETE("", app.DeleteBucket)
	}

	object := router.Group("/:bucket")
	{
		object.PUT("/*key", app.PutObject)
		object.POST("", app.PostObject)
		object.HEAD("/*key", app.HeadObject)
		object.GET("/*key", app.GetObject)
		object.DELETE("/*key", app.DeleteObject)
	}

	service := router.Group("/")
	{
		service.GET("", app.Service)
	}

}
