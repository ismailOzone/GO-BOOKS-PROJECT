// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/ismailOzone/GO-BOOKS-PROJECT/pkg/books/models"
	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// Createbook provides a mock function with given fields: ctx, book
func (_m *Store) Createbook(ctx *context.Context, book *models.Book) error {
	ret := _m.Called(ctx, book)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, *models.Book) error); ok {
		r0 = rf(ctx, book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Deletebook provides a mock function with given fields: ctx, id
func (_m *Store) Deletebook(ctx *context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetbookByID provides a mock function with given fields: ctx, id
func (_m *Store) GetbookByID(ctx *context.Context, id string) (*models.Book, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(*context.Context, string) (*models.Book, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(*context.Context, string) *models.Book); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(*context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Getbooks provides a mock function with given fields: ctx
func (_m *Store) Getbooks(ctx *context.Context) ([]*models.Book, error) {
	ret := _m.Called(ctx)

	var r0 []*models.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(*context.Context) ([]*models.Book, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(*context.Context) []*models.Book); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(*context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Updatebook provides a mock function with given fields: ctx, book
func (_m *Store) Updatebook(ctx *context.Context, book *models.Book) error {
	ret := _m.Called(ctx, book)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, *models.Book) error); ok {
		r0 = rf(ctx, book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}