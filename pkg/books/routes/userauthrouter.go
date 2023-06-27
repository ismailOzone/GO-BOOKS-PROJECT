package routes

import (
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/user/handlers"
	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", handlers.Signup)
	incomingRoutes.POST("users/login", handlers.Login)
}