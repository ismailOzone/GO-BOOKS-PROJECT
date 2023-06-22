package handlers

// import (
// 	// "github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/models"
// 	"net/http"
//     "github.com/gin-gonic/gin"
// 	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/services"
// )

// func Signup(c *gin.Context){
	  
// 	   var body struct{
//         Password string
// 		Email string
//     }
//     if err := c.Bind(&body); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{
//             "error": "Failed to read body",
//         })
//         return
//     }

// 	result , err := services.SignupService(body.Email, body.Password)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Failed to create user",
// 			"cause": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{})
// }