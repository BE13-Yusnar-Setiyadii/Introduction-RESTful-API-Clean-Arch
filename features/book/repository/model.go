package repository

import (
	_book "yusnar/clean-arch/features/book"
	"yusnar/clean-arch/features/user"
	_user "yusnar/clean-arch/features/user/repository"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Publisher   string
	Author      string
	PublishYear string
	UserID      uint
	User        _user.UserForMapping
}

func fromCore(dataCore _book.Core) Book {
	userGorm := Book{
		Title:       dataCore.Title,
		Publisher:   dataCore.Publisher,
		Author:      dataCore.Author,
		PublishYear: dataCore.PublishYear,
		UserID:      dataCore.UserID,
	}
	return userGorm
}
func (dataModel *Book) toCore() _book.Core {
	return _book.Core{
		ID:          dataModel.ID,
		Title:       dataModel.Title,
		Publisher:   dataModel.Publisher,
		Author:      dataModel.Author,
		PublishYear: dataModel.PublishYear,
		UserID:      dataModel.UserID,
		User: user.CoreForMapping{
			Name:    dataModel.User.Name,
			Email:   dataModel.User.Email,
			Address: dataModel.User.Address,
			Role:    dataModel.User.Role,
		},
	}

}

func toCoreList(models []Book) []_book.Core {
	var bookCore []_book.Core
	for _, v := range models {
		bookCore = append(bookCore, v.toCore())

	}
	return bookCore
}

// func toModelList(core []_book.Core) []Book {
// 	var model []Book
// 	for _, v := range core {
// 		model = append(model, fromCore(v))

// 	}
// 	return model

// }
