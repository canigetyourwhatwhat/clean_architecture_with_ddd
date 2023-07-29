package api

import (
	"clean_architecture_with_ddd/internal/controller/entity/request"
	"clean_architecture_with_ddd/internal/controller/inputPort"
	"clean_architecture_with_ddd/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type userHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) UserHandler {
	return &userHandler{
		usecase: u,
	}
}

type UserHandler interface {
	RegisterUser(c echo.Context) error
}

func (uh *userHandler) RegisterUser(c echo.Context) error {
	// Retrieve input
	var body request.CreateUser
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind the struct with the request body: "+err.Error())
	}

	// validate input
	if err := inputPort.CreateUser(body.FirstName, body.LastName, body.Password, body.Username); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// pass to usecase
	if err := uh.usecase.CreateUser(body.FirstName, body.LastName, body.Password, body.Username); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "user created")
}

//type Server interface {
//	AuthHandler
//	CartHandler
//	CartItemHandler
//	OrderHandler
//	PaymentHandler
//	ProductHandler
//	UserHandler
//}

type Server struct {
	ah  AuthHandler
	ch  CartHandler
	cih CartItemHandler
	oh  OrderHandler
	pah PaymentHandler
	prh ProductHandler
	uh  UserHandler
}

//type TestServer struct {
//	ah AuthHandler
//}
//
//type TestInterface interface {
//	Login(c echo.Context) error
//}
//
//func NewTestServer(ah AuthHandler) TestInterface {
//	return TestServer{
//		ah: ah,
//	}
//}
//
//func NewServer(ah AuthHandler, ch CartHandler, cih CartItemHandler, oh OrderHandler, pah PaymentHandler, prh ProductHandler, uh UserHandler) Server {
//	return Server{
//		ah:  ah,
//		ch:  ch,
//		cih: cih,
//		oh:  oh,
//		pah: pah,
//		prh: prh,
//		uh:  uh,
//	}
//}
