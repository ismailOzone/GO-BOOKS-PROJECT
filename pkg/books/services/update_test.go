package services

import (
	"context"
	"errors"
	"testing"

	mocks "github.com/ismailOzone/GO-BOOKS-PROJECT/mocks"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestInsertItem(t *testing.T) {
	type args struct {
		data models.Book
	}

	testCases := []struct {
		name    string
		args    args
		srvc    *BookService
		wantErr error
	}{
		{
			name: "insertion success",
			args: args{
				data: models.Book{
					ID:       "id",
					Name:     "name",
					Author:   "author",
					Year:     2023,
					Language: "language",
				},
			},
			srvc: func() *BookService {
				store := new(mocks.Store)
				// store.On("Updatebook", "BookName", "AuthorName").Return(models.Book{},errors.New("Item not found"))
				
				store.On("Updatebook", mock.Anything, mock.AnythingOfType("*models.Book")).Return(nil)

				return &BookService{
					store: store,
				}
			}(),
			wantErr: nil,
		},
		{
			name: "Updation failure",
			args: args{
				data: models.Book{
					ID:       "id",
					Name:     "name",
					Author:   "author",
					Year:      2023,
					Language: "language",
				},
			},
			srvc: func() *BookService {
				store := new(mocks.Store)
				// store.On("Updatebook", "BookName", "AuthorName").Return(models.Book{}, errors.New("Item not found"))
				store.On("Updatebook", mock.Anything, mock.AnythingOfType("*models.Book")).Return(errors.New("database error"))
				return &BookService{
					store: store,
				}
			}(),
			wantErr: errors.New("database error"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctx:=context.Background()
			err := tt.srvc.Updatebook(&ctx,&tt.args.data)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}
