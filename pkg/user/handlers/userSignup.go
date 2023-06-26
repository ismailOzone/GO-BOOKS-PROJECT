package handlers

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"net/http"

// 	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
// 	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/common/database"

// 	"github.com/gin-gonic/gin"
// 	// "github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/services"
// 	"golang.org/x/crypto/bcrypt"
// )

// func Signup(c *gin.Context) error {
	  
// 	   var body struct{
//         Password string
// 		Email string
//     }
//     if err := c.Bind(&body); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{
//             "error": "Failed to read body",
//         })
//         return err
//     }

// 	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return err
// 	}

// 	user := models.User{
// 		Password: string(hash),
// 		Email:    body.Email,
// 	}
	
// 	db := database.Get()
// 	var buf bytes.Buffer
// 	ctx:=context.Background()
// 	if err := json.NewEncoder(&buf).Encode(user); err != nil {
// 		return err
// 	}
// 	result := db.Index(&ctx , buf , books , user.ID )
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Failed to create user",
// 			"cause": err.Error(),
// 		})
// 		return err
// 	}

// 	c.JSON(http.StatusOK, gin.H{})
// 	return nil
// }