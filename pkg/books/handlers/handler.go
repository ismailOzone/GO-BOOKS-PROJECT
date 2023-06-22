package handlers

import "github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/services"

type BookHandler struct {
	service services.Service
}

var bookHandler *BookHandler

func New() *BookHandler {
	if bookHandler == nil {
		bookHandler = &BookHandler{service: services.New()}
	}
	return bookHandler
}