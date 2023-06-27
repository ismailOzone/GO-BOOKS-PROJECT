package store

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	// "log"

	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/common/database"
	"github.com/pkg/errors"
	// "github.com/elastic/go-elasticsearch/v7"
)

type BookStore struct {
	elasticClient database.ElasticDB
	index         string
}

type Store interface {
	Getbooks(ctx *context.Context) ([]*models.Book, error)
	GetbookByID(ctx *context.Context,id string) (*models.Book, error)
	Createbook(ctx *context.Context,book *models.Book) error
	Updatebook(ctx *context.Context, book *models.Book) error
	Deletebook(ctx *context.Context, id string) error
}

var bookStore Store

// NewBookStore creates a new instance of BookStore
func New() Store {
	if bookStore == nil {
		bookStore = &BookStore{
		elasticClient: database.Get(),
		index:         "books",
		}
	}
	return bookStore
}



func (s *BookStore) Createbook(c *context.Context,book *models.Book) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(book); err != nil {
		return err
	}
	
	err := s.elasticClient.Index(c, buf, s.index, book.ID)
	return err
}

func (s *BookStore) Updatebook(c *context.Context,book *models.Book) error {

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(book); err != nil {
		return err
	}

	err := s.elasticClient.Index(c, buf, s.index, book.ID)
	return err
}


func (s *BookStore) Deletebook(c *context.Context, bookid string) error {
	err := s.elasticClient.Delete(c, bookid, s.index)
	return errors.Wrap(err,"error deleting from es")
}

func (s *BookStore) Getbooks(c *context.Context) ([]*models.Book, error) {
	var buf bytes.Buffer
	ctx:=context.Background()
	searchResult, err := s.elasticClient.Search(&ctx, buf , s.index)
	if err != nil {
		return nil, errors.Wrap(err,"errror doing search from es")
	}

	var response map[string]interface{}
	if err := json.NewDecoder(searchResult).Decode(&response); err != nil {
		return nil, errors.Wrap(err,"errror decoding result from es")
	}

	hits := response["hits"].(map[string]interface{})["hits"].([]interface{})

	var books []*models.Book
	for _, hit := range hits {
		id:=hit.(map[string]interface{})["_id"]
		source := hit.(map[string]interface{})["_source"]
		bookJSON, err := json.Marshal(source)
		if err != nil {
			return nil, errors.Wrap(err,"errror doing marshall from es")
		}
		var book models.Book
		if err := json.Unmarshal(bookJSON, &book); err != nil {
			return nil, errors.Wrap(err,"errror doing json unmarshall from es")
		}
		book.ID,_=id.(string)
		books = append(books, &book)
	}

	return books, nil
}


func (s *BookStore) GetbookByID(c *context.Context, bookID string) (*models.Book, error) {
	should := make([]interface{}, 0, 1)
	should = append(should, map[string]interface{}{
		"match": map[string]interface{}{
			"_id": bookID,
		},
	})

	query := map[string]interface{}{
		"query": should[0],
	}

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	resBody, err := s.elasticClient.Search(c, buf, s.index)
	if err != nil {
		return nil, err
	}

	// log.Printf("********", resBody)

	var hits struct {
		Hits struct {
			Total struct {
				Value int64 `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source *models.Book `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(resBody).Decode(&hits); err != nil {
		return nil, err
	}

	if len(hits.Hits.Hits) == 0 {
		return nil, fmt.Errorf("book not found")
	}

	return hits.Hits.Hits[0].Source, nil
}

