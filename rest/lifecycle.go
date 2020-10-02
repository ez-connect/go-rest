package rest

import (
	"github.com/ez-connect/go-rest/db"
	"github.com/labstack/echo/v4"
)

type LifeCycle interface {
	BeforeFind(c echo.Context, filter *interface{}, option *db.FindOption, projection interface{}) error
	AfterFind(c echo.Context, err error, docs interface{}) error
	BeforeFindOne(c echo.Context, filter, projection interface{}) error
	AfterFindOne(c echo.Context, err error, doc interface{}) error
	BeforeInsert(c echo.Context, doc interface{}) error
	AfterInsert(c echo.Context, err error, doc interface{}) error
	BeforeUpdateOne(c echo.Context, filter, doc interface{}) error
	AfterUpdateOne(c echo.Context, err error, doc interface{}) error
	BeforeDeleteOne(c echo.Context, filter interface{}) error
	AfterDeleteOne(c echo.Context, err error, doc interface{}) error
}
