package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
)

// Createbook godoc
// @Summary      Create a Book
// @Description  Endpoint to create a Book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        Books body models.Book true "Book"
// @Success 201 {string}  "Book created successfully"
// @Failure 404 {object}   object "Book could not be created"
// @Failure 500 {object}   gin.H
// @Router       /books [post]
func (h *BookHandler) Createbook(c *gin.Context) {
	var bookData models.Book
	if err := c.BindJSON(&bookData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		c.Abort()
		return
	}
	ctx:= context.Background()
	err := h.service.Createbook(&ctx, &bookData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create books entry"})
		return
	}

	c.JSON(http.StatusOK, &bookData)
}
