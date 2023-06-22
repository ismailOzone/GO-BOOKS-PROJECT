package services

import (
	"context"

	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/store"
)

type Service interface {
	Getbooks(c *context.Context) ([]*models.Book, error)
	GetbookByID(c *context.Context,id string) (*models.Book, error)
	Createbook(c *context.Context,book *models.Book) error
	Updatebook(c *context.Context,book *models.Book) error
	Deletebook(c *context.Context,id string) error
}

type BookService struct {
	store store.Store
}

var bookService *BookService

func New() *BookService {
	if bookService == nil {
		bookService = &BookService{store: store.New()}
	}
	return bookService
}
