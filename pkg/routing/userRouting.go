package routing

import (
	"github.com/eslamward/helpdesk/pkg/auth"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouting(userServices *auth.UserServices, router *gin.Engine) {
	router.POST("user/register", userServices.RegisterUser)
	router.POST("user/login", userServices.Login)
	router.POST("user/logout", userServices.Logout)
	router.PATCH("user/reset", userServices.ResetPassword)
}
