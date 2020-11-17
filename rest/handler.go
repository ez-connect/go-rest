package rest

import (
	"net/http"
	"strconv"

	"github.com/ez-connect/go-rest/db"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
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

func (h *HandlerBase) Find(c echo.Context,
	filter interface{}, option db.FindOption,
	projection, docs interface{}) error {

	err, total := h.repo.Find(filter, option, projection, docs)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Response().Header().Set(HeaderTotalCount, strconv.Itoa(int(total)))
	return c.JSON(http.StatusOK, docs)
}

func (h *HandlerBase) FindOne(c echo.Context,
	filter, projection interface{}, doc interface{}) error {

	if err := h.repo.FindOne(filter, projection, doc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, doc)
}

func (h *HandlerBase) Aggregate(c echo.Context,
	pipeline, docs interface{}) error {

	total, err := h.repo.Aggregate(pipeline, docs)
	if err == nil {
		c.Response().Header().Set(HeaderTotalCount, strconv.Itoa(int(total)))
		return c.JSON(http.StatusOK, docs)
	}

	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func (h *HandlerBase) AggregateOne(c echo.Context,
	pipeline interface{}, doc interface{}) error {

	err := h.repo.AggregateOne(pipeline, doc)
	if err == nil {
		return c.JSON(http.StatusOK, doc)
	}

	return echo.NewHTTPError(http.StatusNotFound, err.Error())
}

/// Find one document without body
func (h *HandlerBase) Head(c echo.Context, filter interface{}) error {
	count := h.repo.Head(filter)
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

	// Vaildate on insert only
	if err := c.Validate(doc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.repo.Insert(doc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *HandlerBase) UpdateOne(c echo.Context,
	filter, doc interface{}) error {

	if err := Bind(c, doc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.repo.UpdateOne(filter, bson.M{"$set": doc})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h *HandlerBase) DeleteOne(c echo.Context, filter interface{}) error {
	res, err := h.repo.DeleteOne(filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
