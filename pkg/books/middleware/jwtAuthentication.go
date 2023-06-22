package middleware

// import (
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
// 	// "github.com/spf13/afero/internal/common"
// 	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/common/database"
// )
// func RequiredAuth( c *gin.Context){
//     // Get the cookie off request
//     tokenString := c.GetHeader("Authorization")
//     if tokenString == "" {
//         c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
//     }
//     // validate
//         //parse takes the token string and a function for looking up the key
//     token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//         if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//             return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//         }
//         return ([]byte("SECRET")), nil
//     })
//     if err != nil {
//         c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
//     }
//     if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
//         // check the expiry
//         if float64(time.Now().Unix()) > claims["exp"].(float64) {
//             c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
//         }
//         // Find the user with token
//         var user models.User
//         config.DB.First(&user, claims["sub"])
// 		database.ConnectElasticsearch()
//         if user.ID == 0{
//             c.AbortWithStatus(http.StatusUnauthorized)
//         }
//        // attach to req
//        c.Set("user", user)
//        //Continue
//        c.Next()
//     } else {
//         c.AbortWithStatus(http.StatusUnauthorized)
//     }
// }