package service

import (
	"errors"
	"testing"
	"yusnar/clean-arch/features/user"
	"yusnar/clean-arch/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.UserRepository)
	returnData := []user.Core{{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty", Telp_number: "081234", Address: "Jakarta", Role: "user"}}
	//test for success
	t.Run("success get all user", func(t *testing.T) {
		repo.On("GetAll").Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		repo.AssertExpectations(t)
	})
	//test for failed
	t.Run("failed get all user", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestInsert(t *testing.T) {
	repo := new(mocks.UserRepository)
	//test for success
	t.Run("success insert data", func(t *testing.T) {
		insertData := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Telp_number: "081234", Address: "Jakarta", Role: "user"}
		repo.On("Insert", insertData).Return(nil).Once()
		srv := New(repo)
		err := srv.Insert(insertData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	//test for failed
	t.Run("failed insert data, duplicate", func(t *testing.T) {
		insertData := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Telp_number: "081234", Address: "Jakarta", Role: "user"}
		repo.On("Insert", insertData).Return(errors.New("failed insert data")).Once()
		srv := New(repo)
		err := srv.Insert(insertData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("failed insert data, ", func(t *testing.T) {
		insertData := user.Core{Role: ""}
		srv := New(repo)
		err := srv.Insert(insertData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.UserRepository)
	returnData := user.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty", Telp_number: "081234", Address: "Jakarta", Role: "user"}
	//test for success
	t.Run("success get user by id", func(t *testing.T) {
		repo.On("GetById", 1).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, returnData.Name, response.Name)
		repo.AssertExpectations(t)
	})
	//test for failed
	t.Run("failed get user by id", func(t *testing.T) {
		repo.On("GetById", 1).Return(user.Core{}, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.Nil(t, err)
		assert.NotEqual(t, returnData.Name, response.Name)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.UserRepository)
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
	repo := new(mocks.UserRepository)
	//test for success
	t.Run("success update data", func(t *testing.T) {
		insertData := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Telp_number: "081234", Address: "Jakarta", Role: "user"}
		repo.On("Update", 1, insertData).Return(nil).Once()
		srv := New(repo)
		err := srv.Update(1, insertData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	//test for failed
	t.Run("failed update data", func(t *testing.T) {
		insertData := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Telp_number: "081234", Address: "Jakarta", Role: "user"}
		repo.On("Update", 1, insertData).Return(errors.New("failed")).Once()
		srv := New(repo)
		err := srv.Update(1, insertData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

}
