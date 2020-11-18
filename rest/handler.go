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

	err, total := h.repo.Find(params, f, o, projection, docs)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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

	if err := h.repo.FindOne(params, f, projection, doc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, doc)
}

func (h *HandlerBase) Aggregate(c echo.Context,
	pipeline, docs interface{}) error {

	params := filter.GetRawParams(c)
	total, err := h.repo.Aggregate(params, pipeline, docs)
	if err == nil {
		c.Response().Header().Set(HeaderTotalCount, strconv.Itoa(int(total)))
		return c.JSON(http.StatusOK, docs)
	}

	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func (h *HandlerBase) AggregateOne(c echo.Context,
	pipeline interface{}, doc interface{}) error {

	params := filter.GetRawParams(c)
	err := h.repo.AggregateOne(params, pipeline, doc)
	if err == nil {
		return c.JSON(http.StatusOK, doc)
	}

	return echo.NewHTTPError(http.StatusNotFound, err.Error())
}

/// Find one document without body
func (h *HandlerBase) Head(c echo.Context, f interface{}) error {
	params := filter.GetRawParams(c)
	count := h.repo.Head(params, f)
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
	res, err := h.repo.Insert(params, doc, c.Validate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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

	res, err := h.repo.UpdateOne(params, f, doc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h *HandlerBase) DeleteOne(c echo.Context, doc interface{}) error {

	params := filter.GetRawParams(c)
	f, err := filter.FindOne(params, doc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.repo.DeleteOne(params, f)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
