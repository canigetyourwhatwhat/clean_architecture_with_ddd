package api

import (
	"clean_architecture_with_ddd/internal/controller/entity/request"
	"clean_architecture_with_ddd/internal/controller/inputPort"
	"clean_architecture_with_ddd/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) UserHandler {
	return UserHandler{
		usecase: u,
	}
}

func (uh *UserHandler) RegisterUser(c echo.Context) error {
	// Retrieve input
	var body request.CreateUser
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind the struct with the request body: "+err.Error())
	}

	// validate input
	if err := inputPort.CreateUser(body.FirstName, body.LastName, body.Password, body.Username); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// pass to usecase
	if err := uh.usecase.CreateUser(body.FirstName, body.LastName, body.Password, body.Username); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "user created")
}
