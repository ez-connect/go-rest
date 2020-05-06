package rest

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

// Use this binding instead of Echo.Binding()
func Bind(c echo.Context, doc interface{}) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, doc)
}
