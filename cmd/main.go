package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/routes"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// @title Go + Gin + Gorm User API
// @version 1.0
// @description This is a sample Rest API server.
// @host localhost:9000
// @BasePath /\
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Viper
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	// Get the server address port from the config file
	port := viper.GetString("SERVER_ADDRESS")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	log.Println("Starting Init")
	routes.AuthenticationRoutes(router)
	routes.BookRoutes(router)
	log.Println("Started server")
	router.Run(":" + port)
}
