package api

import (
	"clean_architecture_with_ddd/internal/controller/entity/request"
	"clean_architecture_with_ddd/internal/controller/middleware"
	"clean_architecture_with_ddd/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type paymentHandler struct {
	usecase usecase.PaymentUsecase
	auth    middleware.Auth
}

func NewPaymentHandler(u usecase.PaymentUsecase, authorization middleware.Auth) PaymentHandler {
	return &paymentHandler{
		usecase: u,
		auth:    authorization,
	}
}

type PaymentHandler interface {
	CreatePayment(c echo.Context) error
}

func (ph *paymentHandler) CreatePayment(c echo.Context) error {
	// Retrieve input
	userId, err := ph.auth.ValidateSession(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var body request.CreatePayment
	if err = c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind the struct with the request body: "+err.Error())
	}

	// pass to usecase
	err = ph.usecase.CompleteShopping(userId, int(body.Method))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "payment is completed")
}
