package api

import (
	"clean_architecture_with_ddd/internal/controller/entity/request"
	"clean_architecture_with_ddd/internal/controller/inputPort"
	"clean_architecture_with_ddd/internal/controller/middleware"
	"clean_architecture_with_ddd/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CartItemHandler struct {
	usecase usecase.CartItemUsecase
	auth    middleware.Auth
}

func NewCartItemHandler(u usecase.CartItemUsecase, auth middleware.Auth) CartItemHandler {
	return CartItemHandler{
		usecase: u,
		auth:    auth,
	}
}

func (cih *CartItemHandler) AddItemToCart(c echo.Context) error {
	// Retrieve input
	userId, err := cih.auth.GetSession(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var body request.CartItem
	if err = c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind the struct with the request body: "+err.Error())
	}

	// validate input
	if err = inputPort.CartItem(body.ProductCode, body.Quantity); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// pass to usecase
	err = cih.usecase.AddItemInCart(userId, body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "product is added")
}

func (cih *CartItemHandler) RemoveItemFromCart(c echo.Context) error {
	// Retrieve input
	userId, err := cih.auth.GetSession(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var body request.CartItem
	if err = c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind the struct with the request body: "+err.Error())
	}

	// validate input
	if err = inputPort.ProductCode(body.ProductCode); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// pass to usecase
	err = cih.usecase.DeleteItemFromCart(userId, body.ProductCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "product is removed")
}

func (cih *CartItemHandler) UpdateCart(c echo.Context) error {
	// Retrieve input
	userId, err := cih.auth.GetSession(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var body request.ListCartItem
	if err = c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind the struct with the request body: "+err.Error())
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

func (cih *CartItemHandler) GetPurchasedProducts(c echo.Context) error {
	// Retrieve input
	userId, err := cih.auth.GetSession(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// pass to usecase
	cartItems, err := cih.usecase.GetPurchasedProducts(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, cartItems)
}
