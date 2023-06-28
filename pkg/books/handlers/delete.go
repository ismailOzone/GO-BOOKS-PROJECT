package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Deletebook godoc
// @Summary Delete a book
// @Description Deletes a book with the specified ID
// @Tags Books
// @Accept       json
// @Produce      json
// @Param id path int true "Book ID"
// @Success 200 {object} object
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Router /book/{id} [delete]
func (h *BookHandler) Deletebook(c *gin.Context) {
	id := c.Param("id")
	ctx:=c.Request.Context()
	err := h.service.Deletebook(&ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book entry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}