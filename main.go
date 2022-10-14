package main

import (
	"final-project-ems/config"
	"final-project-ems/handler"
	"final-project-ems/repository"
	"final-project-ems/routes"
	"final-project-ems/usecase"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "final-project-ems/docs"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// @title Final Project ECommerce System API Documentation
// @description This is the documentation for the Final Project ECommerce System API
// @version 1.0.0
// @BasePath /api/v1
// @contact.name "Umar Sahid"
// @contact.email "umarhshid@gmail.com"

func main() {
	config.Database()

	config.Migrate()

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Validator = &CustomValidator{validator: validator.New()}

	adminRepository := repository.NewAdminRepository(config.DB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	adminHandler := handler.NewAdminHandler(adminUseCase)

	productRepository := repository.NewProductRepository(config.DB)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productHandler := handler.NewProductHandler(productUseCase)

	productCategoryRepository := repository.NewProductCategoryRepository(config.DB)
	productCategoryUseCase := usecase.NewProductCategoryUseCase(productCategoryRepository)
	productCategoryHandler := handler.NewProductCategoryHandler(productCategoryUseCase)

	userRepository := repository.NewUserRepository(config.DB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	cartRepository := repository.NewCartRepository(config.DB)
	cartUseCase := usecase.NewCartUseCase(cartRepository)
	cartHandler := handler.NewCartHandler(cartUseCase, productUseCase)

	checkoutRepository := repository.NewCheckoutRepository(config.DB)
	checkoutUseCase := usecase.NewCheckoutUseCase(checkoutRepository)
	checkoutHandler := handler.NewCheckoutHandler(checkoutUseCase, cartUseCase)

	routes.Routes(e, adminHandler, productCategoryHandler, userHandler, productHandler, cartHandler, checkoutHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
