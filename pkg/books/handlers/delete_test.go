package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ismailOzone/GO-BOOKS-PROJECT/mocks"

	// "github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeletebook(t *testing.T) {
	testCases := []struct {
		desc          string
		service       *mocks.Service
		bookID        string
		expStatusCode int
	}{
		{
			desc: "success",
			service: func() *mocks.Service {
				mockService := new(mocks.Service)
				mockService.On("Deletebook", mock.Anything,"bookID").Return(nil)
				return mockService
			}(),
			bookID:        "bookID",
			expStatusCode: http.StatusOK,
		},
		{
			desc: "failure - service error",
			service: func() *mocks.Service {
				mockService := new(mocks.Service)
				mockService.On("Deletebook", mock.Anything,"bookID").Return(errors.New("delete error"))
				return mockService
			}(),
			bookID:        "bookID",
			expStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			handler := &BookHandler{
				service: tC.service,
			}
			server := gin.Default()
			server.DELETE("/api/books/:id", handler.Deletebook)

			request := httptest.NewRequest(http.MethodDelete, "/api/books/"+tC.bookID, nil)
			recorder := httptest.NewRecorder()
			server.ServeHTTP(recorder, request)

			assert.Equal(t, tC.expStatusCode, recorder.Code)
		})
	}
}
