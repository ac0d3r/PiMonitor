package middleware

import (
	"pimonitor/database"
	"pimonitor/handler/base"
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			token = c.Query("token")
		} else {
			token = strings.TrimPrefix(authorization, "Bearer ")
		}

		if token == "" {
			base.ReturnUnauthorized(c, 2000)
			c.Abort()
			return
		}

		user, err := database.FindUserByToken(token)
		if err != nil {
			base.ReturnUnauthorized(c, 2000, err)
			c.Abort()
			return
		}

		c.Set("auth.user_id", user.ID)
		c.Set("auth.username", user.Username)
		c.Next()
	}
}
