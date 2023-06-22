package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/routes"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("SERVER_ADDRESS")
	if port==""{
		port="8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	//routes.AuthenticationRoutes(router)
	log.Println("Starting Init")
	routes.BookRoutes(router)
	log.Println("Started server")
	router.Run( port)
    
}