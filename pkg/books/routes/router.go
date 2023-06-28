package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/handlers"
)

func BookRoutes(br *gin.Engine) {
	// incomingRoutes.Use(middleware.RequiredAuth)
	h := handlers.New()

	bookGroup := br.Group("/books")
	{
		bookGroup.GET("", h.Getbooks)
		bookGroup.GET("/:id", h.GetbookByID)
		bookGroup.POST("", h.Createbook)
		bookGroup.PUT("/:id", h.UpdateBookByID)
		bookGroup.DELETE("/:id", h.Deletebook)
	}
}
