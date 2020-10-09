package filter

import (
	"strconv"
	"strings"

	"github.com/ez-connect/go-rest/db"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func Find(c echo.Context, v interface{}) (bson.M, error) {
	queryParams, err := GetQueryParam(c, v)
	if err != nil {
		return nil, err
	}

	pathParams, err := GetPathParam(c, v)
	if err != nil {
		return nil, err
	}

	if pathParams != nil && queryParams != nil {
		return bson.M{"$and": []bson.M{pathParams, queryParams}}, nil
	} else if pathParams != nil {
		return pathParams, nil
	} else if queryParams != nil {
		return queryParams, nil
	} else {
		return bson.M{}, nil
	}
}

func FindOne(c echo.Context, v interface{}) (bson.M, error) {
	return GetPathParam(c, v)
}

func GetQueryParam(c echo.Context, v interface{}) (bson.M, error) {
	query := c.QueryParam("filter")
	if query != "" {
		return UnmarshalQueryParam(query, v)
	} else {
		params := map[string]string{}
		for paramName, param := range c.QueryParams() {
			if len(param) > 0 && !strings.HasPrefix(paramName, "_") {
				params[paramName] = param[0]
			}
		}

		return UnmarshalPathParams(params, v)
	}
}

func GetPathParam(c echo.Context, v interface{}) (bson.M, error) {
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
