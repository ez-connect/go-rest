package filter

import (
	"strconv"
	"strings"

	"github.com/ez-connect/go-rest/db"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type Params struct {
	QueryParams map[string]string
	PathParams  map[string]string
}

func Find(params Params, v interface{}) (bson.M, error) {
	queryParams, err := GetQueryParam(params, v)
	if err != nil {
		return nil, err
	}

	pathParams, err := GetPathParam(params, v)
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

func FindOne(params Params, v interface{}) (bson.M, error) {
	return GetPathParam(params, v)
}

func GetQueryParam(params Params, v interface{}) (bson.M, error) {
	query := params.QueryParams["filter"]
	if query != "" {
		return UnmarshalQueryParam(query, v)
	} else {
		p := map[string]string{}
		for paramName, param := range params.QueryParams {
			// skip option and text search params
			if strings.HasPrefix(paramName, "_") || paramName == "q" {
				continue
			}

			// skip empty param
			if len(param) > 0 {
				p[paramName] = param
			}
		}

		return UnmarshalPathParams(p, v)
	}
}

func GetPathParam(params Params, v interface{}) (bson.M, error) {
	return UnmarshalPathParams(params.PathParams, v)
}

func Option(params Params) db.FindOption {
	options := db.FindOption{Skip: 0}
	if order := params.QueryParams["_order"]; order != "" {
		options.Order = strings.ToLower(order)
	}
	if sort := params.QueryParams["_sort"]; sort != "" {
		options.Sort = sort
	}
	if start, err := strconv.Atoi(params.QueryParams["_start"]); err == nil {
		options.Skip = int64(start)
	}
	if end, err := strconv.Atoi(params.QueryParams["_end"]); err == nil {
		options.Limit = int64(end) - options.Skip
	}
	return options
}

func GetRawParams(c echo.Context) Params {
	params := Params{
		PathParams:  map[string]string{},
		QueryParams: map[string]string{},
	}

	// path params
	for _, paramName := range c.ParamNames() {
		params.PathParams[paramName] = c.Param(paramName)
	}

	// query params
	for paramName, param := range c.QueryParams() {
		// skip empty param
		if len(param) > 0 {
			params.QueryParams[paramName] = param[0]
		}
	}

	return params
}
