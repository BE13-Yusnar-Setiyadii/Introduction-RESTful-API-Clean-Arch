package repository

import (
	"errors"
	"yusnar/clean-arch/features/book"

	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) book.RepositoryInterface {
	return &bookRepository{
		db: db,
	}
}

func (repo *bookRepository) GetAll() ([]book.Core, error) {
	var book []Book
	tx := repo.db.Preload("User").Find(&book)
	if tx.Error != nil {
		return nil, tx.Error
	}
	res := toCoreList(book)
	return res, nil

}

func (repo *bookRepository) Insert(input book.Core) error {
	bookGorm := fromCore(input)
	tx := repo.db.Create(&bookGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil

}

func (repo *bookRepository) GetById(id int) (book.Core, error) {
	books := Book{}
	tx := repo.db.Preload("User").First(&books, id)
	if tx.Error != nil {
		return book.Core{}, tx.Error
	}
	result := books.toCore()
	return result, nil
}

func (repo *bookRepository) Delete(id int) error {
	books := Book{}
	tx := repo.db.Delete(&books, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")

	}
	return nil

}

func (repo *bookRepository) Update(id int, input book.Core) error {
	book := fromCore(input)
	tx := repo.db.Model(book).Where("id = ?", id).Updates(book)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
