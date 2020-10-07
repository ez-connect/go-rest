package rest

import (
	"github.com/ez-connect/go-rest/db"
	"github.com/labstack/echo/v4"
)

type LifeCycle struct {
	BeforeFind      func(c echo.Context, filter *interface{}, option *db.FindOption, projection interface{}) error
	AfterFind       func(c echo.Context, err error, docs interface{}) error
	BeforeFindOne   func(c echo.Context, filter, projection interface{}) error
	AfterFindOne    func(c echo.Context, err error, doc interface{}) error
	BeforeInsert    func(c echo.Context, doc interface{}) error
	AfterInsert     func(c echo.Context, err error, doc interface{}) error
	BeforeUpdateOne func(c echo.Context, filter, doc interface{}) error
	AfterUpdateOne  func(c echo.Context, err error, doc interface{}) error
	BeforeDeleteOne func(c echo.Context, filter interface{}) error
	AfterDeleteOne  func(c echo.Context, err error, doc interface{}) error
}
