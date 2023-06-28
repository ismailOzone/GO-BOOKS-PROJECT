package services

import (
	"context"
	"log"
)


func (s *BookService) Deletebook(c *context.Context ,id string) error {
	
	err := s.store.Deletebook(c , id)

	if err != nil {
		log.Println("delete store eror")
		return err
	}
	return nil
}