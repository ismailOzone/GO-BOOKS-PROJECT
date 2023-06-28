package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
)

// Getbooks godoc
// @Summary Get a list of books
// @Description Retrieves a list of books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Failure 500 {object} object
// @Failure 401 {object} object
// @Router /books [get]
func (h *BookHandler) Getbooks(c *gin.Context) {
	var book []*models.Book

	ctx:= context.Background()
	book, err := h.service.Getbooks(&ctx)
	if err != nil {
		log.Println(err,"error listing books")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}

	c.JSON(http.StatusOK, book)
}