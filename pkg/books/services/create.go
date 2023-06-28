package services

import (
	"context"

	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
) 



func (s *BookService) Createbook(c *context.Context , book *models.Book) error {
	err := s.store.Createbook(c, book)
	if err != nil {
		return err
	}
	return nil
}
