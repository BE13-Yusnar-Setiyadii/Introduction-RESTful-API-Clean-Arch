package repository

import (
	"errors"
	"yusnar/clean-arch/features/user"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.RepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) GetAll() ([]user.Core, error) {
	var user []User
	tx := repo.db.Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	res := toCoreList(user)
	return res, nil

}

func (repo *userRepository) Insert(input user.Core) error {
	userGorm := fromCore(input)
	tx := repo.db.Create(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil

}

func (repo *userRepository) GetById(id int) (user.Core, error) {
	users := User{}
	tx := repo.db.First(&users, id)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	result := users.toCore()
	return result, nil
}

func (repo *userRepository) Delete(id int) error {
	users := User{}
	tx := repo.db.Delete(&users, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")

	}
	return nil

}

func (repo *userRepository) Update(id int, input user.Core) error {
	user := fromCore(input)
	tx := repo.db.Model(user).Where("id = ?", id).Updates(user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
