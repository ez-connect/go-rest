package rest

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

///////////////////////////////////////////////////////////////////

var apiKey string

func InitKeyAuthMiddleware(key string) error {
	if key == "" {
		return errors.New("API Key not found")
	}

	apiKey = key
	return nil
}

func KeyAuthMiddleware() echo.MiddlewareFunc {
	return middleware.KeyAuth(keyAuthValidator)
}

func keyAuthValidator(key string, c echo.Context) (bool, error) {
	return key == apiKey, nil
}
