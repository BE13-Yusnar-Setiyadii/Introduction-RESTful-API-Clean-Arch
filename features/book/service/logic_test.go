package service

import (
	"errors"
	"testing"
	"yusnar/clean-arch/features/book"
	"yusnar/clean-arch/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.BookRepository)
	returnData := []book.Core{{ID: 1, Title: "alta", Publisher: "alta@mail.id", Author: "qwerty", PublishYear: "081234", UserID: 1}}
	//test for success
	t.Run("success get all book", func(t *testing.T) {
		repo.On("GetAll").Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Title, response[0].Title)
		repo.AssertExpectations(t)
	})
	//test for failed
	t.Run("failed get all book", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestInsert(t *testing.T) {
	repo := new(mocks.BookRepository)
	//test for success
	t.Run("success insert data", func(t *testing.T) {
		insertData := book.Core{Title: "alta", Publisher: "alta@mail.id", Author: "qwerty", PublishYear: "081234", UserID: 1}
		repo.On("Insert", insertData).Return(nil).Once()
		srv := New(repo)
		err := srv.Insert(insertData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	//test for failed
	t.Run("failed insert data, duplicate", func(t *testing.T) {
		insertData := book.Core{Title: "alta", Publisher: "alta@mail.id", Author: "qwerty", PublishYear: "081234", UserID: 1}
		repo.On("Insert", insertData).Return(errors.New("failed insert data")).Once()
		srv := New(repo)
		err := srv.Insert(insertData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	// t.Run("failed insert data, ", func(t *testing.T) {
	// 	insertData := book.Core{Role: ""}
	// 	srv := New(repo)
	// 	err := srv.Insert(insertData)
	// 	assert.NotNil(t, err)
	// 	repo.AssertExpectations(t)
	// })
}

func TestGetById(t *testing.T) {
	repo := new(mocks.BookRepository)
	returnData := book.Core{ID: 1, Title: "alta", Publisher: "alta@mail.id", Author: "qwerty", PublishYear: "081234", UserID: 1}
	//test for success
	t.Run("success get book by id", func(t *testing.T) {
		repo.On("GetById", 1).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, returnData.Title, response.Title)
		repo.AssertExpectations(t)
	})
	//test for failed
	t.Run("failed get book by id", func(t *testing.T) {
		repo.On("GetById", 1).Return(book.Core{}, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.Nil(t, err)
		assert.NotEqual(t, returnData.Title, response.Title)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.BookRepository)
	//test for success
	t.Run("success delete data", func(t *testing.T) {
		repo.On("Delete", 1).Return(nil).Once()
		srv := New(repo)
		err := srv.Delete(1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	//test for failed
	t.Run("failed delete data", func(t *testing.T) {
		repo.On("Delete", 1).Return(errors.New("failed")).Once()
		srv := New(repo)
		err := srv.Delete(1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.BookRepository)
	//test for success
	t.Run("success update data", func(t *testing.T) {
		insertData := book.Core{Title: "alta", Publisher: "alta@mail.id", Author: "qwerty", PublishYear: "081234", UserID: 1}
		repo.On("Update", 1, insertData).Return(nil).Once()
		srv := New(repo)
		err := srv.Update(1, insertData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	//test for failed
	t.Run("failed update data", func(t *testing.T) {
		insertData := book.Core{Title: "alta", Publisher: "alta@mail.id", Author: "qwerty", PublishYear: "081234", UserID: 1}
		repo.On("Update", 1, insertData).Return(errors.New("failed")).Once()
		srv := New(repo)
		err := srv.Update(1, insertData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

}
