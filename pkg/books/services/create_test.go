package services

import (
	"context"
	"errors"
	"testing"

	"github.com/ismailOzone/GO-BOOKS-PROJECT/mocks"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatebook(t *testing.T) {
	type args struct {
		data models.Book
	}

	testCases := []struct {
		name    string
		args    args
		srvc    *BookService
		wantErr error
	}{
		// {
		// 	name: "book already exists",
		// 	args: args{
		// 		data: models.Book{
		// 			ID:       "BookID",
		// 			Name:     "BookName",
		// 			Author:   "AuthorName",
		// 			Year:     2023,
		// 			Language: "BookLanguage",
		// 		},
		// 	},
		// 	srvc: func() *BookService {
		// 		store := new(mocks.Store)
		// 		store.On("Createbook", "BookName", "AuthorName").Return(models.Book{}, nil)

		// 		return &BookService{
		// 			store: store,
		// 		}
		// 	}(),
		// 	wantErr: errors.New("Book already exists"),
		// },
		{
			name: "book creation success",
			args: args{
				data: models.Book{
					ID:       "BookID",
					Name:     "BookName",
					Author:   "AuthorName",
					Year:     2023,
					Language: "BookLanguage",
				},
			},
			srvc: func() *BookService {
				store := new(mocks.Store)
				// store.On("GetBookByNameAndAuthor", "BookName", "AuthorName").Return(models.Book{}, errors.New("Book not found"))
				store.On("Createbook", mock.Anything,mock.Anything).Return(nil)

				return &BookService{
					store: store,
				}
			}(),
			wantErr: nil,
		},
		{
			name: "book creation failure",
			args: args{
				data: models.Book{
					ID:       "BookID",
					Name:     "BookName",
					Author:   "AuthorName",
					Year:     2023,
					Language: "BookLanguage",
				},
			},
			srvc: func() *BookService {
				store := new(mocks.Store)
				// store.On("GetBookByNameAndAuthor", "BookName", "AuthorName").Return(models.Book{}, errors.New("Book not found"))
				store.On("Createbook", mock.Anything,mock.Anything).Return(errors.New("database error"))

				return &BookService{
					store: store,
				}
			}(),
			wantErr: errors.New("database error"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			err := tt.srvc.Createbook(&ctx, &tt.args.data)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}
