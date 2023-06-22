package models


type Book struct {
	Name    string  `json:"name"`
	Author  string  `json:"author"`
	Year    uint16   `json:"year"`
	Language string `json:"language"`
	ID       string    `json:"id"`
}


