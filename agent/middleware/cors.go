package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
)

func Cors() gin.HandlerFunc {
	return cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "OPTIONS, GET, PUT, POST, DELETE, HEAD",
		RequestHeaders:  "x-kss-newfilename-in-body, Origin, Authorization, Content-Type,  X-Requested-With",
		ExposedHeaders:  "newfilename, x-kss-request-id, ETag, Date",
		MaxAge:          10 * time.Minute,
		Credentials:     false,
		ValidateHeaders: true,
	})

}
