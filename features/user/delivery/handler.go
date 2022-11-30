package delivery

import (
	"log"
	"net/http"
	"strconv"
	"yusnar/clean-arch/features/user"
	"yusnar/clean-arch/middlewares"
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
	e.GET("/users", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/users/:id", handler.GetID, middlewares.JWTMiddleware())
	e.POST("/users", handler.Create)
	e.PUT("/users/:id", handler.UpdateID, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", handler.DeleteId, middlewares.JWTMiddleware())
}

func (delivery *UserDelivery) GetAll(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idToken:", idToken)

	roleToken := middlewares.ExtractTokenUserRole(c)
	log.Println("roleToken:", roleToken)

	result, err := delivery.userServices.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed"))
	}
	data := responseList(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get All Users", data))

}

func (delivery *UserDelivery) Create(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idToken:", idToken)

	roleToken := middlewares.ExtractTokenUserRole(c)
	log.Println("roleToken:", roleToken)

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
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Create user"))

}

func (delivery *UserDelivery) GetID(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idToken:", idToken)

	roleToken := middlewares.ExtractTokenUserRole(c)
	log.Println("roleToken:", roleToken)

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
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get user", result))

}

func (delivery *UserDelivery) DeleteId(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idToken:", idToken)

	roleToken := middlewares.ExtractTokenUserRole(c)
	log.Println("roleToken:", roleToken)

	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}
	err := delivery.userServices.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Delete"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Delete user"))

}

func (delivery *UserDelivery) UpdateID(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idToken:", idToken)

	roleToken := middlewares.ExtractTokenUserRole(c)
	log.Println("roleToken:", roleToken)

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
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Update user"))

}
