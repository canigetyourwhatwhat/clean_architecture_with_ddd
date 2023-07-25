package api

import (
	"clean_architecture_with_ddd/internal/controller/entity/request"
	"clean_architecture_with_ddd/internal/controller/inputPort"
	"clean_architecture_with_ddd/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(u usecase.AuthUsecase) AuthHandler {
	return AuthHandler{
		usecase: u,
	}
}

func (ah *AuthHandler) Login(c echo.Context) error {
	// Retrieve input
	var body request.LoginInput
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind the struct with the request body: "+err.Error())
	}

	// validate input
	if err := inputPort.Login(body.Username, body.Password); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// pass to usecase
	session, err := ah.usecase.Login(body.Username, body.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, session)
}
