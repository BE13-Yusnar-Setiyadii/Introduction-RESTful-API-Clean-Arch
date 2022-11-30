package service

import (
	"yusnar/clean-arch/features/auth"
)

type authService struct {
	authRepository auth.RepositoryInterface
}

func NewAuth(repo auth.RepositoryInterface) auth.ServiceInterface {
	return &authService{
		authRepository: repo,
	}
}

func (service *authService) Login(email string, pass string) (string, error) {
	data, err := service.authRepository.Login(email, pass)
	return data, err

}
