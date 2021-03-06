package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func WithTransaction(callback func(sessCtx mongo.SessionContext) (interface{}, error), opts ...*options.TransactionOptions) (interface{}, error) {
	ctx := context.Background()
	session, err := GetMongoDb().GetClient().(*mongo.Client).StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(ctx)

	result, err := session.WithTransaction(ctx, callback, opts...)
	if err != nil {
		return nil, err
	}
	fmt.Printf("result: %v\n", result)
	return result, nil
}
