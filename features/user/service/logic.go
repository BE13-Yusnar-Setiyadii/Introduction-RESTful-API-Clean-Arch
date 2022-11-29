package service

import (
	"yusnar/clean-arch/features/user"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userRepository user.RepositoryInterface
	validate       *validator.Validate
}

func New(repo user.RepositoryInterface) user.ServiceInterface {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

func (service *userService) GetAll() (data []user.Core, err error) {
	data, err = service.userRepository.GetAll()
	return data, err
}

func (service *userService) Insert(input user.Core) error {
	input.Role = "user"
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	err := service.userRepository.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

func (service *userService) GetById(id int) (user.Core, error) {
	data, errData := service.userRepository.GetById(id)
	if errData != nil {
		return user.Core{}, nil
	}
	return data, nil

}

func (service *userService) Delete(id int) error {
	errDelete := service.userRepository.Delete(id)
	if errDelete != nil {
		return errDelete
	}
	return nil
}

func (service *userService) Update(id int, input user.Core) error {
	err := service.userRepository.Update(id, input)
	if err != nil {
		return err
	}
	return nil
}
