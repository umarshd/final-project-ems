package routes

import (
	"final-project-ems/handler"
	"final-project-ems/middlewares"

	"github.com/labstack/echo/v4"
)

func Routes(
	e *echo.Echo,
	adminHandler *handler.AdminHandler,
	productCategoryHandler *handler.ProductCategoryHandler,
	userHandler *handler.UserHandler,
	productHandler *handler.ProductHandler,
	cartHandler *handler.CartHandler,
	checkoutHandler *handler.CheckoutHandler,
) {

	e.POST("/admin/register", adminHandler.CreateAdmin)
	e.POST("/admin/login", adminHandler.LoginAdmin)

	e.POST("/products", productHandler.CreateProduct, middlewares.AuthenticationAdmin())
	e.GET("/products", productHandler.FindAllProduct)
	e.GET("/products/:id", productHandler.FindByIdProduct)
	e.POST("/products/filter", productHandler.FindProduct)

	e.POST("/product-category", productCategoryHandler.CreateProductCategory, middlewares.AuthenticationAdmin())
	e.GET("/product-category", productCategoryHandler.GetAllProductCategory, middlewares.AuthenticationAdmin())

	e.POST("/users/register", userHandler.CreateUser)
	e.POST("/users/login", userHandler.LoginUser)

	e.POST("users/:id/address", userHandler.CreateUserAddress, middlewares.AuthenticationUser())

	e.POST("/carts", cartHandler.CreateCart, middlewares.AuthenticationUser())
	e.GET("/carts", cartHandler.FindAllCart, middlewares.AuthenticationUser())

	e.POST("/cart-items", cartHandler.CreateCartItem, middlewares.AuthenticationUser())

	e.POST("/checkout", checkoutHandler.CreateCheckout, middlewares.AuthenticationUser())
}
