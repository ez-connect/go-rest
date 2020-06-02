package db

import (
	"go.mongodb.org/mongo-driver/bson"
)

type DatabaseBase interface {
	Init(config interface{}) error
	Connect()
	Close()
	GetClient() interface{}
	GetCursor(collection string,
		filter, sort interface{}, skip, limit int64,
		projection interface{}) (interface{}, error)

	Find(collection string,
		filter interface{}, option FindOption,
		projection, docs interface{}) error

	Distinct(collection, fieldName string,
		filter interface{}) ([]string, error)

	Aggregate(collection string, pipeline interface{},
		docs interface{}) error

	FindOne(collection string, filter interface{},
		projection interface{}, doc interface{}) error

	AggregateOne(collection string, pipeline interface{},
		doc interface{}) error

	// Insert executes an insert command to insert a single document into the collection.
	// The _id can be retrieved from the result.
	Insert(collection string, doc interface{}) (interface{}, error)
	InsertMany(collection string, docs []interface{}) ([]interface{}, error)
	UpdateOne(collection string, filter interface{},
		doc interface{}) (interface{}, error)
	UpdateMany(collection string, filter interface{},
		doc interface{}) (interface{}, error)

	DeleteOne(collection string, filter interface{}) (interface{}, error)
	DeleteMany(collection string, filter interface{}) (interface{}, error)
	Count(collection string, filter interface{}) (int64, error)

	EnsureIndex(collection string, name string, keys bson.M,
		unique bool) string
}
