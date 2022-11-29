package delivery

import (
	"net/http"
	"strconv"
	"yusnar/clean-arch/features/user"
	"yusnar/clean-arch/utils/helper"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userServices user.ServiceInterface
}

func New(service user.ServiceInterface, e *echo.Echo) {
	handler := &UserDelivery{
		userServices: service,
	}
	e.GET("/users", handler.GetAll)
	e.GET("/users/:id", handler.GetID)
	e.POST("/users", handler.Create)
	e.PUT("/users/:id", handler.UpdateID)

	e.DELETE("/users/:id", handler.DeleteId)
}

func (delivery *UserDelivery) GetAll(c echo.Context) error {
	result, err := delivery.userServices.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed"))
	}
	data := responseList(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get All Users", data))

}

func (delivery *UserDelivery) Create(c echo.Context) error {
	user := UserRequest{}
	errBind := c.Bind(&user)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}

	result := user.reqToCore()

	err := delivery.userServices.Insert(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error insert"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Get Create user"))

}

func (delivery *UserDelivery) GetID(c echo.Context) error {
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}

	userId, err := delivery.userServices.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delivery"))
	}

	result := coreToResponse(userId)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get Create user", result))

}

func (delivery *UserDelivery) DeleteId(c echo.Context) error {
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}
	err := delivery.userServices.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Delete"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Delete"))

}

func (delivery *UserDelivery) UpdateID(c echo.Context) error {
	user := UserRequest{}
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}

	errBind := c.Bind(&user)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}

	result := user.reqToCore()
	err := delivery.userServices.Update(id, result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Update"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Update"))

}
