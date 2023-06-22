package services

import (
	"context"
	"errors"

	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
)

func (s *BookService) Updatebook(c *context.Context , book *models.Book) error {
	if book == nil {
		return errors.New("book is required")
	}
	return s.store.Updatebook(c , book)
}