package repository

import (
	"errors"
	"yusnar/clean-arch/features/auth"
	"yusnar/clean-arch/features/user/repository"
	"yusnar/clean-arch/middlewares"

	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuth(db *gorm.DB) auth.RepositoryInterface {
	return &authRepository{
		db: db,
	}
}

func (repo *authRepository) Login(email string, pass string) (string, error) {
	var userData repository.User
	tx := repo.db.Where("email = ? AND password = ?", email, pass).First(&userData)
	if tx.Error != nil {
		return "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return "", errors.New("login failed")
	}

	token, errToken := middlewares.CreateToken(int(userData.ID), userData.Role)
	if errToken != nil {
		return "", errToken
	}

	return token, nil
}
