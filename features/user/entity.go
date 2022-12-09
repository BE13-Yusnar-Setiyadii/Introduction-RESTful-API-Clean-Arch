package user

import (
	"time"
)

type Core struct {
	ID          uint
	Name        string `validate:"required"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required"`
	Telp_number string `validate:"required"`
	Address     string
	Role        string `validate:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// Book        []BookCoreForMapping
}

type CoreForMapping struct {
	Name    string
	Email   string
	Address string
	Role    string
}

// type BookCoreForMapping struct {
// 	Title       string
// 	Publisher   string
// 	Author      string
// 	PublishYear string
// }

type ServiceInterface interface {
	GetAll() (data []Core, err error)
	Insert(input Core) error
	GetById(id int) (Core, error)
	Delete(id int) error
	Update(id int, input Core) error
}

type RepositoryInterface interface {
	GetAll() ([]Core, error)
	Insert(input Core) error
	GetById(id int) (Core, error)
	Delete(id int) error
	Update(id int, input Core) error
}
