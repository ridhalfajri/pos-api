package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ridhalfajri/pos-api/internal/app/controller"
)

func SetupRouter(
	router *gin.Engine,
	userController controller.UserController,
) {
	api := router.Group("/api/v1")
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)
	}
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"success": false, "message": "Endpoint tidak ditemukan"})
	})
}
