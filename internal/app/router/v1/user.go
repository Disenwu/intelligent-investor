package router

import (
	"intelligent-investor/internal/app/handler"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	api := r.Group("/v1/user")
	{
		api.POST("/login", handler.LoginHandler)
		api.PUT("/register", handler.RegisterHandler)
	}
}
