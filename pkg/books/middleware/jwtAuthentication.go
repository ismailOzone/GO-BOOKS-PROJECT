package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"

	// "log"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/config"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"

	// "github.com/spf13/afero/internal/common"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/common/database"
)

var esClient *elasticsearch.Client


func RequiredAuth( c *gin.Context){
    // Get the cookie off request
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
    }

	cfg:= config.Get()

    // validate
        //parse takes the token string and a function for looking up the key
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return  []byte(cfg.SecretKey), nil
    })
    if err != nil {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
    }
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
        // check the expiry
        if float64(time.Now().Unix()) > claims["exp"].(float64) {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
        }
        // Find the user with token
        var user models.User
		database.ConnectElasticsearch()

		// Construct the search request
		searchRequest := map[string]interface{}{
			"query": map[string]interface{}{
				"term": map[string]interface{}{
					"id": claims["sub"],
				},
			},
		}

		// Encode the search request as JSON
		buf := new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(searchRequest)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		res, err := esClient.Search(
			esClient.Search.WithIndex("users"),
			esClient.Search.WithBody(buf),
		)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer res.Body.Close()

			// Check if there are any hits/user exists
			if res.StatusCode != http.StatusOK {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
	
			// Parse the response body
			var searchResponse map[string]interface{}
			err = json.NewDecoder(res.Body).Decode(&searchResponse)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
	
			hits := searchResponse["hits"].(map[string]interface{})["hits"].([]interface{})
			if len(hits) == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

       // attach to req
       c.Set("user", user)
       //Continue
       c.Next()
    } else {
        c.AbortWithStatus(http.StatusUnauthorized)
    }
}