package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
)

// UpdateBookByID godoc
// @Summary Update a book
// @Description Updates an existing book
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Book object"
// @Success 200 {object} object
// @Failure 400 {object} object
// @Failure 404 {object} object
// @Failure 401 {object} object
// @Router /book/{id} [put]
func (h *BookHandler) UpdateBookByID(c *gin.Context) {
	id := c.Param("id")


	var bookData *models.Book
    if c.BindJSON(&bookData) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "details are not provided properly"})
		c.Abort()
		return
	} 

	bookData.ID = id
	ctx:= context.Background()
	err := h.service.Updatebook( &ctx,bookData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book entry"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}