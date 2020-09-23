package filter

import (
	"github.com/ez-connect/go-rest/db"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func Find(c echo.Context) bson.M {
	return nil
}

func Option(c echo.Context) db.FindOption {
	return db.FindOption{}
}

func FindOne(c echo.Context) bson.M {
	return nil
}
