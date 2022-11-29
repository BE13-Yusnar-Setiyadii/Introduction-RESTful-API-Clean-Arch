package factory

import (
	_delivery "yusnar/clean-arch/features/user/delivery"
	_repository "yusnar/clean-arch/features/user/repository"
	_service "yusnar/clean-arch/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(db *gorm.DB, e *echo.Echo) {
	userRepoFactroy := _repository.New(db)
	userServiceFactory := _service.New(userRepoFactroy)
	_delivery.New(userServiceFactory, e)
}
