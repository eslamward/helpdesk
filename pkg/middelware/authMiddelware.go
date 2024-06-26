package middelware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddelware() gin.HandlerFunc {

	return func(context *gin.Context) {

		session := sessions.Default(context)

		getUserEmail := session.Get("userEmail")

		if getUserEmail == nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"login": "http://localhost/user/login"})

		}
		userEmail := getUserEmail.(string)

		if userEmail == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"login": "http://localhost/user/login"})

		}

		session.Set("userEmail", userEmail)
		context.Next()
	}
}
