package rest

import (
	"github.com/ez-connect/go-rest/db"
	"github.com/labstack/echo/v4"
)

type RouterBase interface {
	Init(e *echo.Echo, db db.DatabaseBase)
}
