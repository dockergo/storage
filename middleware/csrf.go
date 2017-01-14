package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func Csrf() gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret: "storagesecret",
		ErrorFunc: func(c *gin.Context) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "CSRF token mismatch",
			})
			c.Abort()
		},
	})
}
