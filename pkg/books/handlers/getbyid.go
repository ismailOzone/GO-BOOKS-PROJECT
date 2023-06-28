package handlers

import (
	"context"
	"net/http"

	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
)

// GetbookByID godoc
// @Summary Get a book by ID
// @Description Retrieves a book by its ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} object
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 404 {object} object
// @Router /book/{id} [get]
func (h *BookHandler) GetbookByID(c *gin.Context) {
	id := c.Param("id")
	ctx:= context.Background()
	var book *models.Book
	var err error
	book, err = h.service.GetbookByID(&ctx , id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get book by ID"})
		return
	}

	c.JSON(http.StatusOK, book)
}