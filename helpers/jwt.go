package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "RahasiaBanget"

func GenerateTokenAdmin(id string, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"admin": true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))

	return signedToken, err
}

func GenerateTokenUser(id string, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"admin": false,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))

	return signedToken, err
}

func ValidateToken(authHeader string) (interface{}, error) {
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	fmt.Println(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errors.New("Invalid token")
	} else {
		return claims, nil
	}
}
