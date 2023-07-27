package cmd

import (
	"clean_architecture_with_ddd/internal/controller/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(ah api.AuthHandler, ch api.CartHandler, cih api.CartItemHandler, oh api.OrderHandler, pah api.PaymentHandler, prh api.ProductHandler, uh api.UserHandler) *echo.Echo {

	e := echo.New()

	// resolving CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:9000"},
		AllowMethods:     []string{"CREAT", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	e.POST("/register", uh.RegisterUser)
	e.POST("/login", ah.Login)

	product := e.Group("/product")
	product.GET("list", prh.ListProducts)
	product.GET(":code", prh.GetProductByCode)

	cart := e.Group("/cart")
	cart.GET("", ch.GetInProgressCart)

	cartItem := e.Group("/cartItem")
	cartItem.POST("", cih.AddItemToCart)
	cartItem.DELETE(":code", cih.RemoveItemFromCart)
	cartItem.PUT("", cih.UpdateCart)
	cartItem.GET("", cih.GetPurchasedProducts)

	payment := e.Group("/payment")
	payment.POST("", pah.CreatePayment)

	order := e.Group("/order")
	order.GET(":id", oh.GetOrder)

	return e
}
