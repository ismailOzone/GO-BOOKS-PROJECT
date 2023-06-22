package services

// import (
// 	"github.com/elastic/go-elasticsearch"
// 	"github.com/ismailOzone/GO-BOOKS-PROJECT/config"
// 	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/models"
// 	"golang.org/x/crypto/bcrypt"
// )

// // SignupService performs the user signup and returns the result.
// func SignupService(email string, password string) error {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
// 	if err != nil {
// 		return err
// 	}

// 	user := models.User{
// 		Password: string(hash),
// 		Email:    email,
// 	}

// 	var client *elasticsearch.Client
//     // result := config.ElasticDB.
// 	result := config.DB.Create(&user)
// 	return result
// }