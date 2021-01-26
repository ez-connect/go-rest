package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type DatabaseBase interface {
	Init(config interface{}) error
	Connect()
	Close()
	GetClient() interface{}
	GetCollection(collection string) interface{}
	GetCursor(ctx context.Context, collection string,
		filter, sort interface{}, skip, limit int64,
		projection interface{}) (interface{}, error)

	Find(ctx context.Context, collection string,
		filter interface{}, option FindOption,
		projection, docs interface{}) error

	Distinct(ctx context.Context, collection, fieldName string,
		filter interface{}) ([]string, error)

	Aggregate(ctx context.Context, collection string, pipeline interface{},
		docs interface{}) error

	FindOne(ctx context.Context, collection string, filter interface{},
		projection interface{}, doc interface{}) error

	AggregateOne(ctx context.Context, collection string, pipeline interface{},
		doc interface{}) error

	// Insert executes an insert command to insert a single document into the collection.
	// The _id can be retrieved from the result.
	Insert(ctx context.Context, collection string, doc interface{}) (InsertOneResult, error)
	InsertMany(ctx context.Context, collection string, docs []interface{}) ([]interface{}, error)
	UpdateOne(ctx context.Context, collection string, filter, doc interface{}) (UpdateOneResult, error)
	FindOneAndUpdate(ctx context.Context, collection string, filter, update, doc interface{}) error
	UpdateMany(ctx context.Context, collection string, filter interface{},
		doc interface{}) (interface{}, error)

	DeleteOne(ctx context.Context, collection string, filter interface{}) (interface{}, error)
	DeleteMany(ctx context.Context, collection string, filter interface{}) (interface{}, error)
	Count(ctx context.Context, collection string, filter interface{}) (int64, error)

	EnsureIndex(ctx context.Context, collection string, name string, keys bson.M,
		unique bool) string
}

type InsertOneResult struct {
	Id interface{} `json:"id"`
}

type UpdateOneResult struct {
	Id interface{} `json:"id"`
}
