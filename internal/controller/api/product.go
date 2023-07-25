package api

import (
	"clean_architecture_with_ddd/internal/controller/inputPort"
	"clean_architecture_with_ddd/internal/controller/presenter"
	"clean_architecture_with_ddd/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductHandler struct {
	usecase usecase.ProductUsecase
}

func NewProductHandler(u usecase.ProductUsecase) ProductHandler {
	return ProductHandler{
		usecase: u,
	}
}

func (ph *ProductHandler) ListProducts(c echo.Context) error {

	// Retrieve input
	q := c.Request().URL.Query()
	pageNum := q.Get("page")

	// validate input
	page, err := inputPort.ListProductsByPage(pageNum)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// pass to usecase
	products, count, err := ph.usecase.ListProductsByPage(page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// pass to presenter
	res := presenter.AllProduct(count, products)

	return c.JSON(http.StatusOK, res)
}

func (ph *ProductHandler) GetProductByCode(c echo.Context) error {

	// Retrieve input
	code := c.Param("code")

	// validate input
	err := inputPort.ID(code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// pass to usecase
	product, err := ph.usecase.GetProduct(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, product)
}
