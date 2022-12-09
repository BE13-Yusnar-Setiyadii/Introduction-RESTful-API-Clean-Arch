package service

import (
	"yusnar/clean-arch/features/book"

	"github.com/go-playground/validator/v10"
)

type bookService struct {
	bookRepository book.RepositoryInterface
	validate       *validator.Validate
}

func New(repo book.RepositoryInterface) book.ServiceInterface {
	return &bookService{
		bookRepository: repo,
		validate:       validator.New(),
	}
}

func (service *bookService) GetAll() (data []book.Core, err error) {
	data, err = service.bookRepository.GetAll()
	return data, err
}

func (service *bookService) Insert(input book.Core) error {
	err := service.bookRepository.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

func (service *bookService) GetById(id int) (book.Core, error) {
	data, errData := service.bookRepository.GetById(id)
	if errData != nil {
		return book.Core{}, nil
	}
	return data, nil

}

func (service *bookService) Delete(id int) error {
	errDelete := service.bookRepository.Delete(id)
	if errDelete != nil {
		return errDelete
	}
	return nil
}

func (service *bookService) Update(id int, input book.Core) error {
	err := service.bookRepository.Update(id, input)
	if err != nil {
		return err
	}
	return nil
}
