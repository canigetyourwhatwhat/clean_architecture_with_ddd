package api

import (
	"clean_architecture_with_ddd/internal/controller/inputPort"
	"clean_architecture_with_ddd/internal/controller/middleware"
	"clean_architecture_with_ddd/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type orderHandler struct {
	usecase usecase.OrderUsecase
	auth    middleware.Auth
}

func NewOrderHandler(u usecase.OrderUsecase, authorization middleware.Auth) OrderHandler {
	return &orderHandler{
		usecase: u,
		auth:    authorization,
	}
}

type OrderHandler interface {
	GetOrder(c echo.Context) error
}

func (oh *orderHandler) GetOrder(c echo.Context) error {
	// Retrieve input
	idStr := c.Param("id")
	userId, err := oh.auth.ValidateSession(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// validate input
	orderId, err := inputPort.IntID(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// pass to usecase
	order, err := oh.usecase.GetOrder(userId, orderId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, *order)
}
