package rest

import (
	"github.com/ez-connect/go-rest/db"
	"github.com/ez-connect/go-rest/rest/filter"
	"go.mongodb.org/mongo-driver/bson"
)

type RepositoryBase struct {
	RepositoryInterface
	Driver     db.DatabaseBase
	collection string // collection name aka a collection in NOSQL or a table in SQL
	lifeCycle  LifeCycle
}

type RepositoryInterface interface {
	EnsureIndexs()
	Init(db db.DatabaseBase, collection string)
	RegisterLifeCycle(l LifeCycle)
	Find(params filter.Params, filter interface{}, option db.FindOption,
		projection, docs interface{}) (error, int64)
	FindOne(params filter.Params, filter, projection interface{}, doc interface{}) error
	Aggregate(params filter.Params, pipeline, docs interface{}) (int64, error)
	AggregateOne(params filter.Params, pipeline interface{}, doc interface{}) error
	Head(params filter.Params, filter interface{}) int64
	Insert(params filter.Params, doc interface{}) (interface{}, error)
	UpdateOne(params filter.Params, filter, doc interface{}) (interface{}, error)
	DeleteOne(params filter.Params, filter interface{}) (interface{}, error)
}

func (r *RepositoryBase) Init(db db.DatabaseBase, collection string) {
	r.Driver = db
	r.collection = collection
	// h.lifeCycle = interface{}(h).(LifeCycle)
}

func (r *RepositoryBase) RegisterLifeCycle(l LifeCycle) {
	r.lifeCycle = l
}

func (r *RepositoryBase) Find(params filter.Params, filter interface{}, option db.FindOption,
	projection, docs interface{}) (error, int64) {

	if r.lifeCycle.BeforeFind != nil {
		if err := r.lifeCycle.BeforeFind(params, &filter, &option, &projection); err != nil {
			return err, 0
		}
	}

	err := r.Driver.Find(r.collection, filter, option, projection, docs)

	if r.lifeCycle.AfterFind != nil {
		if err := r.lifeCycle.AfterFind(params, docs); err != nil {
			return err, 0
		}
	}

	if err != nil {
		return err, 0
	}

	total, _ := r.Driver.Count(r.collection, filter)
	return nil, total
}

func (r *RepositoryBase) FindOne(params filter.Params, filter, projection interface{}, doc interface{}) error {

	if r.lifeCycle.BeforeFindOne != nil {
		err := r.lifeCycle.BeforeFindOne(params, filter, projection)
		if err != nil {
			return err
		}
	}

	if err := r.Driver.FindOne(r.collection, filter, projection, doc); err != nil {
		return err
	}

	if r.lifeCycle.AfterFindOne != nil {
		if err := r.lifeCycle.AfterFindOne(params, doc); err != nil {
			return err
		}
	}

	return nil
}

func (r *RepositoryBase) Aggregate(params filter.Params, pipeline, docs interface{}) (int64, error) {

	err := r.Driver.Aggregate(r.collection, pipeline, docs)
	if err != nil {
		return 0, err
	}

	total, _ := r.Driver.Count(r.collection, bson.M{})
	return total, nil
}

func (r *RepositoryBase) AggregateOne(params filter.Params, pipeline interface{}, doc interface{}) error {

	err := r.Driver.AggregateOne(r.collection, pipeline, doc)
	if err != nil {
		return err
	}

	return nil
}

/// Find one document without body
func (r *RepositoryBase) Head(params filter.Params, filter interface{}) int64 {
	count, _ := r.Driver.Count(r.collection, filter)
	return count
}

func (r *RepositoryBase) Insert(params filter.Params, doc interface{}) (interface{}, error) {
	if r.lifeCycle.BeforeInsert != nil {
		if err := r.lifeCycle.BeforeInsert(params, doc); err != nil {
			return nil, err
		}
	}

	res, err := r.Driver.Insert(r.collection, doc)
	if err != nil {
		return nil, err
	}

	if r.lifeCycle.AfterInsert != nil {
		if err := r.lifeCycle.AfterInsert(params, doc); err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (r *RepositoryBase) UpdateOne(params filter.Params, filter, doc interface{}) (interface{}, error) {
	if r.lifeCycle.BeforeUpdateOne != nil {
		if err := r.lifeCycle.BeforeUpdateOne(params, filter, doc); err != nil {
			return nil, err
		}
	}

	res, err := r.Driver.UpdateOne(r.collection, filter, bson.M{"$set": doc})
	if err != nil {
		return nil, err
	}

	if r.lifeCycle.AfterUpdateOne != nil {
		if err := r.lifeCycle.AfterUpdateOne(params, doc); err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (r *RepositoryBase) DeleteOne(params filter.Params, filter interface{}) (interface{}, error) {
	if r.lifeCycle.BeforeDeleteOne != nil {
		if err := r.lifeCycle.BeforeDeleteOne(params, filter); err != nil {
			return nil, err
		}
	}

	res, err := r.Driver.DeleteOne(r.collection, filter)
	if err != nil {
		return nil, err
	}

	if r.lifeCycle.AfterDeleteOne != nil {
		if err := r.lifeCycle.AfterDeleteOne(params, res); err != nil {
			return nil, err
		}
	}

	return res, nil
}
