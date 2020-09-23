package filter

import (
	"github.com/ez-connect/go-rest/db"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Find(c echo.Context, v interface{}) bson.M {
	query := c.QueryParam("q")
	if query == "" {
		return nil
	}
	return Unmarshal(query, v)
}

func FindOne(c echo.Context) bson.M {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return nil
	}

	return bson.M{"_id": id}
}

func Option(c echo.Context) db.FindOption {
	return db.FindOption{}
}
