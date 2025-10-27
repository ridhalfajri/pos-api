package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ridhalfajri/pos-api/internal/app/controller"
	"github.com/ridhalfajri/pos-api/internal/app/router"
	"github.com/ridhalfajri/pos-api/internal/app/service"
	"github.com/ridhalfajri/pos-api/internal/config"
	"github.com/ridhalfajri/pos-api/internal/repository"
)

func main() {
	config.ConnectDB()
	db := config.DB
	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository)
	userController := controller.NewUserControllerImpl(userService)

	routerEngine := gin.Default()
	router.SetupRouter(routerEngine, userController)

	log.Println("Starting POS Web API Server on port 8000...")
	if err := routerEngine.Run(":8000"); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}
