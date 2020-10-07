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
	db         db.DatabaseBase
	collection string // collection name aka a collection in NOSQL or a table in SQL
	lifeCycle  LifeCycle
}

///////////////////////////////////////////////////////////////////

func (h *HandlerBase) Init(db db.DatabaseBase, collection string) {
	h.db = db
	h.collection = collection
	// h.lifeCycle = interface{}(h).(LifeCycle)
}

func (h *HandlerBase) RegisterLifeCycle(l LifeCycle) {
	h.lifeCycle = l
}

func (h *HandlerBase) Find(c echo.Context,
	filter interface{}, option db.FindOption,
	projection, docs interface{}) error {

	// err := h.LifeCycle.BeforeFind(c, &filter, &option, &projection)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusServiceUnavailable, err.Error())
	// }

	err := h.db.Find(h.collection, filter, option, projection, docs)

	// err = h.LifeCycle.AfterFind(c, err, docs)

	if err == nil {
		total, _ := h.db.Count(h.collection, filter)
		c.Response().Header().Set(HeaderTotalCount, strconv.Itoa(int(total)))
		return c.JSON(http.StatusOK, docs)
	}

	return echo.NewHTTPError(http.StatusServiceUnavailable, err.Error())
}

func (h *HandlerBase) FindOne(c echo.Context,
	filter, projection interface{}, doc interface{}) error {

	err := h.lifeCycle.BeforeFindOne(c, filter, projection)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.db.FindOne(h.collection, filter, projection, doc)

	err = h.lifeCycle.AfterFindOne(c, err, doc)

	if err == nil {
		return c.JSON(http.StatusOK, doc)
	}

	c.Logger().Debug(err)
	return echo.NewHTTPError(http.StatusNotFound)
}

func (h *HandlerBase) Aggregate(c echo.Context,
	pipeline, docs interface{}) error {

	err := h.db.Aggregate(h.collection, pipeline, docs)
	if err == nil {
		total, _ := h.db.Count(h.collection, bson.M{})
		c.Response().Header().Set(HeaderTotalCount, strconv.Itoa(int(total)))
		return c.JSON(http.StatusOK, docs)
	}

	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func (h *HandlerBase) AggregateOne(c echo.Context,
	pipeline interface{}, doc interface{}) error {

	err := h.db.AggregateOne(h.collection, pipeline, doc)
	if err == nil {
		return c.JSON(http.StatusOK, doc)
	}

	return echo.NewHTTPError(http.StatusNotFound, err.Error())
}

/// Find one document without body
func (h *HandlerBase) Head(c echo.Context, filter interface{}) error {
	count, _ := h.db.Count(h.collection, filter)
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

	err := h.lifeCycle.BeforeInsert(c, doc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Vaildate on insert only
	if err := c.Validate(doc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.db.Insert(h.collection, doc)

	err = h.lifeCycle.AfterInsert(c, err, doc)

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

	err := h.lifeCycle.BeforeUpdateOne(c, filter, doc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.db.UpdateOne(h.collection, filter, bson.M{"$set": doc})

	err = h.lifeCycle.AfterUpdateOne(c, err, doc)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h *HandlerBase) DeleteOne(c echo.Context, filter interface{}) error {

	err := h.lifeCycle.BeforeDeleteOne(c, filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.db.DeleteOne(h.collection, filter)

	err = h.lifeCycle.AfterDeleteOne(c, err, res)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
