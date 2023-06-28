package handlers

import (
	"bytes"
	"context"
	"encoding/json"

	// "log"
	"net/http"
	"time"

	"github.com/ismailOzone/GO-BOOKS-PROJECT/config"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/common/database"

	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	cfg := config.New()

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

	// Check for the user in the database
	var user models.User
	db := database.Get()
	var buf bytes.Buffer
	ctx := context.Background()
	if err := json.NewEncoder(&buf).Encode(body.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to encode user",
		})
		// log.Printf("sssssssssssss", body.Email)
		return
	}

	indexName := "users"
	res, err := db.Search(&ctx, buf, indexName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to search for user",
		})
		return
	}
	defer res.Close()

	var searchResult map[string]interface{}
	if err := json.NewDecoder(res).Decode(&searchResult); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to decode search result",
		})
		return
	}

	hits := searchResult["hits"].(map[string]interface{})
	totalHits := hits["total"].(map[string]interface{})["value"].(float64)

	// Check if the user exists
	if totalHits == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Retrieve the user data from the search result
	hitsArray := hits["hits"].([]interface{})
	firstHit := hitsArray[0].(map[string]interface{})
	source := firstHit["_source"].(map[string]interface{})
	user.Email = source["Email"].(string)
	user.Password = source["Password"].(string)

	// Check if the password matches
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   user.Email,
		"expires": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenStringVal, err := token.SignedString([]byte(cfg.SecretKey))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenStringVal,
	})
}
