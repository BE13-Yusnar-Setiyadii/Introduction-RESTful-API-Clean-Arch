package delivery

import (
	"log"
	"net/http"
	"strconv"
	"yusnar/clean-arch/features/book"
	"yusnar/clean-arch/middlewares"
	"yusnar/clean-arch/utils/helper"

	"github.com/labstack/echo/v4"
)

type BookDelivery struct {
	bookServices book.ServiceInterface
}

func New(service book.ServiceInterface, e *echo.Echo) {
	handler := &BookDelivery{
		bookServices: service,
	}
	e.GET("/books", handler.GetAll)
	e.GET("/books/:id", handler.GetID)
	e.POST("/books", handler.Create, middlewares.JWTMiddleware())
	e.PUT("/books/:id", handler.UpdateID, middlewares.JWTMiddleware())
	e.DELETE("/books/:id", handler.DeleteId, middlewares.JWTMiddleware())
}

func (delivery *BookDelivery) GetAll(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idToken:", idToken)

	roleToken := middlewares.ExtractTokenUserRole(c)
	log.Println("roleToken:", roleToken)

	result, err := delivery.bookServices.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed"))
	}
	data := responseList(result)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get All Books", data))

}

func (delivery *BookDelivery) Create(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idToken:", idToken)

	roleToken := middlewares.ExtractTokenUserRole(c)
	log.Println("roleToken:", roleToken)

	book := BookRequest{}
	errBind := c.Bind(&book)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}
	result := book.reqToCore()
	err := delivery.bookServices.Insert(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error insert"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Create book"))

}

func (delivery *BookDelivery) GetID(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idToken:", idToken)

	roleToken := middlewares.ExtractTokenUserRole(c)
	log.Println("roleToken:", roleToken)

	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}

	bookId, err := delivery.bookServices.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delivery"))
	}

	result := coreToResponse(bookId)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get book", result))

}

func (delivery *BookDelivery) DeleteId(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idToken:", idToken)

	roleToken := middlewares.ExtractTokenUserRole(c)
	log.Println("roleToken:", roleToken)

	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}
	err := delivery.bookServices.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Delete"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Delete book"))

}

func (delivery *BookDelivery) UpdateID(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	log.Println("idToken:", idToken)

	roleToken := middlewares.ExtractTokenUserRole(c)
	log.Println("roleToken:", roleToken)

	book := BookRequest{}
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Convert"))
	}

	errBind := c.Bind(&book)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}

	result := book.reqToCore()
	err := delivery.bookServices.Update(id, result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Update"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Update book"))

}
