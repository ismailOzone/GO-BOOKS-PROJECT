package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	// "time"

	"github.com/gin-gonic/gin"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/mocks"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdatebookbyID(t *testing.T) {
	testCases := []struct {
		desc          string
		service       *mocks.Service
		requestBody   interface{}
		expStatusCode int
	}{
		{
			desc: "success",
			service: func() *mocks.Service {
				mockService := new(mocks.Service)
				mockService.On("UpdatebookbyID", mock.Anything , mock.AnythingOfType("*models.Book")).Return(nil)
				return mockService
			}(),
			requestBody: models.Book{
				Name: "BookName",
				Author : "AuthorName",
				Year: 2023,
                Language: "BookLanguage",
			},
			expStatusCode: http.StatusOK,
		},
		{
			desc: "failure - service error",
			service: func() *mocks.Service {
				mockService := new(mocks.Service)
				mockService.On("UpdatebookbyID", mock.Anything , mock.AnythingOfType("*models.Book")).Return(errors.New("insert error"))
				return mockService
			}(),
			requestBody: models.Book{
				Name: "BookName",
				Author : "AuthorName",
				Year: 2023,
                Language: "BookLanguage",
			},
			expStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			handler := &BookHandler{
				service : tC.service,
			}
			server := gin.Default()
			server.PUT("/api/books/:id", handler.UpdateBookByID)

			requestBody, err := json.Marshal(tC.requestBody)
			assert.NoError(t, err)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPut, "/api/books/:id", bytes.NewBuffer(requestBody))
			server.ServeHTTP(recorder, request)

			assert.Equal(t, tC.expStatusCode, recorder.Code)
			tC.service.AssertExpectations(t)
		})
	}
}