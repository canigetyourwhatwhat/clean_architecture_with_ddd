package api

import (
	"clean_architecture_with_ddd/internal/controller/entity/request"
	"clean_architecture_with_ddd/internal/controller/inputPort"
	"clean_architecture_with_ddd/internal/controller/middleware"
	"clean_architecture_with_ddd/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type cartItemHandler struct {
	usecase usecase.CartItemUsecase
	auth    middleware.Auth
}

func NewCartItemHandler(u usecase.CartItemUsecase, auth middleware.Auth) CartItemHandler {
	return &cartItemHandler{
		usecase: u,
		auth:    auth,
	}
}

type CartItemHandler interface {
	AddItemToCart(c echo.Context) error
	RemoveItemFromCart(c echo.Context) error
	UpdateCart(c echo.Context) error
	GetPurchasedProducts(c echo.Context) error
}

func (cih *cartItemHandler) UpdateCart(c echo.Context) error {

	// Retrieve input
	var body request.ListCartItem
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind the struct with the request body: "+err.Error())
	}

	// middleware validation
	userId, err := cih.auth.ValidateSession(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// validate input
	if err = inputPort.ListCartItem(body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// pass to usecase
	err = cih.usecase.UpdateItemsInCart(userId, body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "cart is updated")
}

func (cih *cartItemHandler) AddItemToCart(c echo.Context) error {
	// Retrieve input
	userId, err := cih.auth.ValidateSession(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var body request.CartItem
	if err = c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind the struct with the request body: "+err.Error())
	}

	// validate input
	if err = inputPort.CartItem(body.ProductCode, body.Quantity); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// pass to usecase
	err = cih.usecase.AddItemInCart(userId, body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "product is added")
}

func (cih *cartItemHandler) RemoveItemFromCart(c echo.Context) error {
	// Retrieve input
	userId, err := cih.auth.ValidateSession(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	code := c.Param("code")

	// validate input
	if err = inputPort.ProductCode(code); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// pass to usecase
	err = cih.usecase.DeleteItemFromCart(userId, code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "product is removed")
}

func (cih *cartItemHandler) GetPurchasedProducts(c echo.Context) error {
	// Retrieve input
	userId, err := cih.auth.ValidateSession(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// pass to usecase
	cartItems, err := cih.usecase.GetPurchasedProducts(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cartItems)
}
