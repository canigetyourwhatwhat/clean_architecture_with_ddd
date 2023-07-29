package api

import (
	"clean_architecture_with_ddd/internal/controller/middleware"
	"clean_architecture_with_ddd/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type cartHandler struct {
	usecase usecase.CartUsecase
	auth    middleware.Auth
}

func NewCartHandler(u usecase.CartUsecase, authorization middleware.Auth) CartHandler {
	return &cartHandler{
		usecase: u,
		auth:    authorization,
	}
}

type CartHandler interface {
	GetInProgressCart(c echo.Context) error
}

func (ch *cartHandler) GetInProgressCart(c echo.Context) error {
	// Retrieve input
	userId, err := ch.auth.ValidateSession(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// pass to usecase
	cart, err := ch.usecase.GetInProgressCart(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cart)
}
