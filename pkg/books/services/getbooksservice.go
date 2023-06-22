package services

import (
	"context"

	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
	// "github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/store"
)

func (s *BookService) Getbooks(c *context.Context) ([]*models.Book, error) {
	return s.store.Getbooks(c)
}