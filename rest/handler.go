package rest

import (
	"net/http"
	"strconv"

	"github.com/ez-connect/go-rest/db"
	"github.com/ez-connect/go-rest/rest/filter"
	"github.com/labstack/echo/v4"
)

///////////////////////////////////////////////////////////////////

const (
	HeaderTotalCount = "X-Total-Count"
)

type HandlerBase struct {
	repo RepositoryInterface
}

///////////////////////////////////////////////////////////////////

func (h *HandlerBase) Init(db db.DatabaseBase, collection string, repo RepositoryInterface) {
	if repo == nil {
		h.repo = &RepositoryBase{}
	} else {
		h.repo = repo
	}
	h.repo.Init(db, collection)
}

func (h *HandlerBase) Find(c echo.Context, projection, docs interface{}) error {
	params := filter.GetRawParams(c)
	f, err := filter.Find(params, docs)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	o := filter.Option(params)

	total, err := h.repo.Find(params, nil, f, o, projection, docs)
	if err != nil {
		return returnError(c, err)
	}

	c.Response().Header().Set(HeaderTotalCount, strconv.Itoa(int(total)))
	return c.JSON(http.StatusOK, docs)
}

func (h *HandlerBase) FindOne(c echo.Context, projection interface{}, doc interface{}) error {

	params := filter.GetRawParams(c)
	f, err := filter.FindOne(params, doc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.repo.FindOne(params, nil, f, projection, doc); err != nil {
		return returnError(c, err)
	}

	return c.JSON(http.StatusOK, doc)
}

func (h *HandlerBase) Aggregate(c echo.Context,
	pipeline, docs interface{}) error {

	params := filter.GetRawParams(c)
	total, err := h.repo.Aggregate(params, nil, pipeline, docs)
	if err == nil {
		c.Response().Header().Set(HeaderTotalCount, strconv.Itoa(int(total)))
		return c.JSON(http.StatusOK, docs)
	}

	return returnError(c, err)
}

func (h *HandlerBase) AggregateOne(c echo.Context,
	pipeline interface{}, doc interface{}) error {

	params := filter.GetRawParams(c)
	err := h.repo.AggregateOne(params, nil, pipeline, doc)
	if err == nil {
		return c.JSON(http.StatusOK, doc)
	}

	return returnError(c, err)
}

/// Find one document without body
func (h *HandlerBase) Head(c echo.Context, f interface{}) error {
	params := filter.GetRawParams(c)
	count := h.repo.Head(params, nil, f)
	if count > 0 {
		c.Response().WriteHeader(http.StatusOK)
	} else {
		c.Logger().Debug("Not found")
		c.Response().WriteHeader(http.StatusNotFound)
	}

	return nil
}

func (h *HandlerBase) Insert(c echo.Context, doc interface{}) error {
	if err := Bind(c, doc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	params := filter.GetRawParams(c)
	res, err := h.repo.Insert(params, nil, doc, c.Validate)
	if err != nil {
		return returnError(c, err)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *HandlerBase) UpdateOne(c echo.Context, doc interface{}) error {

	if err := Bind(c, doc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	params := filter.GetRawParams(c)
	f, err := filter.FindOne(params, doc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.repo.UpdateOne(params, nil, f, doc, c.Validate)
	if err != nil {
		return returnError(c, err)
	}

	return c.JSON(http.StatusOK, res)
}

func (h *HandlerBase) DeleteOne(c echo.Context, doc interface{}) error {

	params := filter.GetRawParams(c)
	f, err := filter.FindOne(params, doc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.repo.DeleteOne(params, nil, f)
	if err != nil {
		return returnError(c, err)
	}
	return c.JSON(http.StatusOK, res)
}

func returnError(c echo.Context, err error) error {
	if _, ok := err.(rerror); ok {
		return echo.NewHTTPError(err.(rerror).Code(), err.Error())
	}
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}
