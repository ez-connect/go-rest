package filter

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ById(c echo.Context) (interface{}, error) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return nil, err
	}

	return bson.M{"_id": id}, nil
}
