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

// func TestInsertItem(t *testing.T) {
// 	type args struct {
// 		data models.Book
// 	}

// 	testCases := []struct {
// 		name    string
// 		args    args
// 		srvc    *Service
// 		wantErr error
// 	}{
// 		{
// 			name: "item already exists",
// 			args: args{
// 				data: models.Book{
// 					ID: "BookID",
// 					Name: "BookName",
// 					Author : "AuthorName",
// 					Year: 2023,
// 					Language: "BookLanguage",
// 				},
// 			},
// 			srvc: func() *Service {
// 				store := new(mocks.Store)
// 				store.On("Updatebook", mock.Anything, mock.Anything).Return(models.Book{}, nil)
// 				return &Service{
// 					store: store,
// 				}
// 			}(),
// 			wantErr: errors.New("Item Details Already there"),
// 		},
// 		{
// 			name: "insertion success",
// 			args: args{
// 				data: models.Book{
// 					ID: "BookID",
// 					Name: "BookName",
// 					Author : "AuthorName",
// 					Year: 2023,
// 					Language: "BookLanguage",
// 				},
// 			},
// 			srvc: func() *Service {
// 				store := new(mocks.Store)
// 				store.On("GetItemByModelBrand", filter).Return(models.Book{}, serror.NotFoundError("Item not found"))
// 				store.On("InsertItem", mock.Anything).Return(nil)
// 				return &Service{
// 					store: store,
// 				}
// 			}(),
// 			wantErr: nil,
// 		},
// 		{
// 			name: "insertion failure",
// 			args: args{
// 				data: models.Book{
// 					ID: "BookID",
// 					Name: "BookName",
// 					Author : "AuthorName",
// 					Year: 2023,
// 					Language: "BookLanguage",
// 				},
// 			},
// 			srvc: func() *Service {
// 				store := new(mocks.Store)
// 				store.On("GetItemByModelBrand", filter).Return(models.Book{}, serror.NotFoundError("Item not found"))
// 				store.On("InsertItem", mock.Anything).Return(errors.New("database error"))
// 				return &Service{
// 					store: store,
// 				}
// 			}(),
// 			wantErr: errors.New("Error adding the Item"),
// 		},
// 	}

// 	for _, tt := range testCases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := tt.srvc.InsertItem(tt.args.data)
// 			assert.Equal(t, err, tt.wantErr)
// 		})
// 	}
// }


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
			name: "item already exists",
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
				// store.On("GetItemByModelBrand", "BookName", "AuthorName").Return(models.Book{}, nil)
				store.On("Updatebook", mock.Anything, mock.AnythingOfType("*models.Book")).Return(models.Book{}, nil)

				return &BookService{
					store: store,
				}
			}(),
			wantErr: errors.New("Item Details Already there"),
		},
		{
			name: "insertion success",
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
				store.On("Updatebook", "BookName", "AuthorName").Return(models.Book{},errors.New("Item not found"))
				
				store.On("InsertItem", mock.Anything).Return(nil)
				return &BookService{
					store: store,
				}
			}(),
			wantErr: nil,
		},
		{
			name: "insertion failure",
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
				// store.On("Updatebook", "BookName", "AuthorName").Return(models.Book{}, errors.New("Item not found"))
				store.On("Updatebook", mock.Anything, mock.AnythingOfType("*models.Book")).Return(models.Book{}, errors.New("Item not found"))

				store.On("InsertItem", mock.Anything).Return(errors.New("database error"))
				return &BookService{
					store: store,
				}
			}(),
			wantErr: errors.New("Error adding the Item"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// err := tt.srvc.InsertItem(tt.args.data)
			ctx:=context.Background()
			err := tt.srvc.Updatebook(&ctx,&tt.args.data)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}
