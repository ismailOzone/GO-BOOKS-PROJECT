package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/common/database"

	"github.com/gin-gonic/gin"
	// "github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/services"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Password string
		Email    string
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Password: string(hash),
		Email:    body.Email,
	}

	db := database.Get()
	var buf bytes.Buffer
	ctx := context.Background()
	if err := json.NewEncoder(&buf).Encode(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to encode user",
		})
		return
	}

	indexName := "books"

	if err := db.Index(&ctx, buf, indexName, ""); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
			"cause": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Created successfully"})
}