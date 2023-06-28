package services

import (
	"context"

	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
	// "github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/store"
)

func (s *BookService) GetbookByID(c *context.Context  , id string) (*models.Book, error) {
	return s.store.GetbookByID(c , id)
}