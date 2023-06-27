package routes

import (
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/middleware"
	"github.com/gin-gonic/gin"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/handlers"
)

func BookRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.RequiredAuth)
	h := handlers.New()

	bookGroup := incomingRoutes.Group("/books")
	{
		bookGroup.GET("", h.Getbooks)
		bookGroup.GET("/:id", h.GetbookByID)
		bookGroup.POST("", h.Createbook)
		bookGroup.PUT("/:id", h.UpdateBookByID)
		bookGroup.DELETE("/:id", h.Deletebook)
	}
}
