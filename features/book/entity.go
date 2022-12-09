package book

import (
	"time"
	"yusnar/clean-arch/features/user"
)

type Core struct {
	ID          uint
	Title       string
	Publisher   string
	Author      string
	PublishYear string
	UserID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        user.CoreForMapping
}

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
