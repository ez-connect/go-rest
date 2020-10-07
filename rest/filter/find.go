package filter

import (
	"strconv"
	"strings"

	"github.com/ez-connect/go-rest/db"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func Find(c echo.Context, v interface{}) bson.M {
	queryParams := GetQueryParam(c, v)
	pathParams := GetPathParam(c, v)
	if pathParams != nil && queryParams != nil {
		return bson.M{"$and": []bson.M{pathParams, queryParams}}
	} else if pathParams != nil {
		return pathParams
	} else if queryParams != nil {
		return queryParams
	} else {
		return nil
	}
}

func FindOne(c echo.Context, v interface{}) bson.M {
	return GetPathParam(c, v)
}

func GetQueryParam(c echo.Context, v interface{}) bson.M {
	query := c.QueryParam("filter")
	if query == "" {
		return nil
	}
	return UnmarshalQueryParam(query, v)
}

func GetPathParam(c echo.Context, v interface{}) bson.M {
	params := map[string]string{}
	for _, paramName := range c.ParamNames() {
		params[paramName] = c.Param(paramName)
	}

	return UnmarshalPathParams(params, v)
}

func Option(c echo.Context) db.FindOption {
	options := db.FindOption{Skip: 0}
	if order := c.QueryParam("_order"); order != "" {
		options.Order = strings.ToLower(order)
	}
	if sort := c.QueryParam("_sort"); sort != "" {
		options.Sort = sort
	}
	if start, err := strconv.Atoi(c.QueryParam("_start")); err == nil {
		options.Skip = int64(start)
	}
	if end, err := strconv.Atoi(c.QueryParam("_end")); err == nil {
		options.Limit = int64(end) - options.Skip
	}
	return options
}
