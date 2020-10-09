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

	if h.lifeCycle.BeforeFind != nil {
		if err := h.lifeCycle.BeforeFind(c, &filter, &option, &projection); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	err := h.db.Find(h.collection, filter, option, projection, docs)

	if h.lifeCycle.AfterFind != nil {
		if err := h.lifeCycle.BeforeFind(c, &filter, &option, &projection); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	if err == nil {
		total, _ := h.db.Count(h.collection, filter)
		c.Response().Header().Set(HeaderTotalCount, strconv.Itoa(int(total)))
		return c.JSON(http.StatusOK, docs)
	}

	return echo.NewHTTPError(http.StatusServiceUnavailable, err.Error())
}

func (h *HandlerBase) FindOne(c echo.Context,
	filter, projection interface{}, doc interface{}) error {

	if h.lifeCycle.BeforeFindOne != nil {
		err := h.lifeCycle.BeforeFindOne(c, filter, projection)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	if err := h.db.FindOne(h.collection, filter, projection, doc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if h.lifeCycle.AfterFindOne != nil {
		if err := h.lifeCycle.AfterFindOne(c, doc); err != nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}
	}

	return c.JSON(http.StatusOK, doc)
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

	if h.lifeCycle.BeforeInsert != nil {
		if err := h.lifeCycle.BeforeInsert(c, doc); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	// Vaildate on insert only
	if err := c.Validate(doc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.db.Insert(h.collection, doc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if h.lifeCycle.AfterInsert != nil {
		if err := h.lifeCycle.AfterInsert(c, doc); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *HandlerBase) UpdateOne(c echo.Context,
	filter, doc interface{}) error {

	if err := Bind(c, doc); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if h.lifeCycle.BeforeUpdateOne != nil {
		if err := h.lifeCycle.BeforeUpdateOne(c, filter, doc); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	res, err := h.db.UpdateOne(h.collection, filter, bson.M{"$set": doc})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if h.lifeCycle.AfterUpdateOne != nil {
		if err := h.lifeCycle.AfterUpdateOne(c, doc); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	return c.JSON(http.StatusOK, res)
}

func (h *HandlerBase) DeleteOne(c echo.Context, filter interface{}) error {
	if h.lifeCycle.BeforeDeleteOne != nil {
		if err := h.lifeCycle.BeforeDeleteOne(c, filter); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	res, err := h.db.DeleteOne(h.collection, filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if h.lifeCycle.AfterDeleteOne != nil {
		if err := h.lifeCycle.AfterDeleteOne(c, res); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	return c.JSON(http.StatusOK, res)
}
