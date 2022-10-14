package middlewares

import (
	"final-project-ems/entity"
	"final-project-ems/helpers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func AuthenticationAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get token from header
			authHeader := c.Request().Header.Get("Authorization")
			token, err := helpers.ValidateToken(authHeader)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, &entity.ErrorResponse{
					Status:  http.StatusUnauthorized,
					Message: err.Error(),
				})
			}
			// Set user information from token into context
			// fmt.Println(token.(jwt.MapClaims)["admin"])
			if token.(jwt.MapClaims)["admin"] != true {
				return c.JSON(http.StatusUnauthorized, &entity.ErrorResponse{
					Status:  http.StatusUnauthorized,
					Message: "You are not admin",
				})
			}

			c.Set("admin", token)

			return next(c)
		}
	}
}

func AuthenticationUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get token from header
			authHeader := c.Request().Header.Get("Authorization")
			token, err := helpers.ValidateToken(authHeader)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, &entity.ErrorResponse{
					Status:  http.StatusUnauthorized,
					Message: err.Error(),
				})
			}
			// Set user information from token into context
			c.Set("user", token)

			return next(c)
		}
	}
}
