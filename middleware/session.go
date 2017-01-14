package middleware

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Sessions() gin.HandlerFunc {
	store := sessions.NewCookieStore([]byte("secret"))
	return sessions.Sessions("storagesession", store)

}
