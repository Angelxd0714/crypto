package main

import (
	_ "gin-quickstart/docs" // Use your project's module path here
	auth "gin-quickstart/internal/auth/controller"
	"gin-quickstart/internal/auth/services"
	"gin-quickstart/internal/database"
	"gin-quickstart/internal/websocket"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Quickstart API
// @version 1.0
// @description This is a sample server for a gin quickstart application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db := database.DB

	authService := services.NewAuthService(db)
	authController := auth.NewAuthController(authService)

	broker := websocket.NewBroker()
	go broker.Run()

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		authRoutes := v1.Group("/auth")
		{
			authRoutes.POST("/register", authController.Register)
			authRoutes.POST("/login", authController.Login)
		}
		v1.GET("/ws", func(c *gin.Context) {
			broker.ServeWs(c)
		})
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run() // listens on 0.0.0.0:8080 by default
}
