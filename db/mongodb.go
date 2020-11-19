package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

///////////////////////////////////////////////////////////////////

// YAML config for MongoDB
type MongoDBConfig struct {
	Host       string `yaml:"host"`
	Username   string `yaml:"user"`
	Password   string `yaml:"password"`
	AuthSource string `yaml:"authSource"`
	Name       string `yaml:"name"`
}

func (c *MongoDBConfig) IsValid() bool {
	if c.Host == "" || c.Username == "" || c.Password == "" {
		return false
	}

	return true
}

///////////////////////////////////////////////////////////////////

type MongoDb struct {
	DatabaseBase

	config  MongoDBConfig
	session *mongo.Client
}

var mongoOnce sync.Once
var mongodb *MongoDb

func GetMongoDb() *MongoDb {
	mongoOnce.Do(func() {
		mongodb = &MongoDb{}
	})
	return mongodb
}

///////////////////////////////////////////////////////////////////

func (db *MongoDb) Init(i interface{}) error {
	fmt.Println("Init MongoDB")
	config := i.(MongoDBConfig)
	if !config.IsValid() {
		return errors.New("Invalid mongodb config")
	}

	db.config = config
	return nil
}

func (db *MongoDb) Connect() {
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s/?authSource=%s",
		db.config.Username,
		db.config.Password,
		db.config.Host,
		db.config.AuthSource,
	)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Failed to create Mongo client:", err)
	}

	if err := client.Connect(context.Background()); err == nil {
		db.session = client
		fmt.Println("Connected to database:", db.config.Name)
	} else {
		log.Fatal("Failed to establish connection to Mongo server: ", err)
	}
}

func (db *MongoDb) Close() {
	if err := db.session.Disconnect(context.Background()); err != nil {
		fmt.Println(err)
	}
}

// Returns the current session, *mongo.Client
func (db *MongoDb) GetClient() interface{} {
	return db.session
}

func (db *MongoDb) GetCursor(ctx context.Context, collection string,
	filter, sort interface{},
	skip, limit int64,
	projection interface{}) (interface{}, error) {

	if ctx == nil {
		ctx = context.TODO()
	}
	opts := options.Find()
	if sort != nil {
		opts.SetSort(sort)
	}
	if skip > 0 {
		opts.SetSkip(skip)
	}
	if limit > 0 {
		opts.SetLimit(limit)
	}

	if projection != nil {
		opts.SetProjection(projection)
	}

	return db.getCollection(collection).Find(ctx, filter, opts)
}

func (db *MongoDb) Find(ctx context.Context, collection string,
	filter interface{}, option FindOption,
	projection, docs interface{}) error {

	if ctx == nil {
		ctx = context.TODO()
	}

	var sort bson.M
	if option.Sort != "" {
		if option.Order == "desc" {
			sort = bson.M{option.Sort: -1}
		} else {
			sort = bson.M{option.Sort: 1}
		}
	}

	res, err := db.GetCursor(
		ctx,
		collection,
		filter,
		sort,
		option.Skip,
		option.Limit,
		projection,
	)
	if err != nil {
		return err
	}

	cur := res.(*mongo.Cursor) // it a cursor
	defer cur.Close(ctx)
	return cur.All(ctx, docs)
}

func (db *MongoDb) Distinct(ctx context.Context, collection, fieldName string,
	filter interface{}) ([]string, error) {

	if ctx == nil {
		ctx = context.TODO()
	}
	res, err := db.getCollection(collection).Distinct(ctx, fieldName, filter)
	values := make([]string, len(res))
	for i := range res {
		values[i] = res[i].(string)
	}
	return values, err
}

func (db *MongoDb) Aggregate(ctx context.Context, collection string, pipeline interface{},
	docs interface{}) error {

	if ctx == nil {
		ctx = context.TODO()
	}
	cur, err := db.getCollection(collection).Aggregate(ctx, pipeline)
	if err == nil {
		defer cur.Close(ctx)
		return cur.All(ctx, docs)
	}
	return err
}

func (db *MongoDb) FindOne(ctx context.Context, collection string, filter interface{},
	projection interface{}, doc interface{}) error {
	if ctx == nil {
		ctx = context.TODO()
	}
	opts := options.FindOne()
	if projection != nil {
		opts.SetProjection(projection)
	}
	res := db.getCollection(collection).FindOne(ctx, filter, opts)
	return res.Decode(doc)
}

func (db *MongoDb) AggregateOne(ctx context.Context, collection string, pipeline interface{},
	doc interface{}) error {
	if ctx == nil {
		ctx = context.TODO()
	}
	cur, err := db.getCollection(collection).Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	if cur.Next(ctx) {
		return cur.Decode(doc)
	}

	return nil
}

func (db *MongoDb) Insert(ctx context.Context, collection string, doc interface{}) (InsertOneResult, error) {
	if ctx == nil {
		ctx = context.TODO()
	}
	mres, err := db.getCollection(collection).InsertOne(ctx, doc)
	res := InsertOneResult{Id: mres.InsertedID}
	return res, err
}

func (db *MongoDb) InsertMany(ctx context.Context, collection string, docs []interface{}) ([]interface{}, error) {
	if ctx == nil {
		ctx = context.TODO()
	}
	res, err := db.getCollection(collection).InsertMany(ctx, docs)
	if err != nil {
		return nil, err
	}
	return res.InsertedIDs, err
}

func (db *MongoDb) UpdateOne(ctx context.Context, collection string, filter interface{},
	update interface{}) (UpdateOneResult, error) {
	if ctx == nil {
		ctx = context.TODO()
	}
	_, err := db.getCollection(collection).UpdateOne(ctx, filter, update)
	var id interface{}
	f, ok := filter.(primitive.M)
	if ok {
		id = f["_id"]
	}
	res := UpdateOneResult{Id: id}
	return res, err
}

func (db *MongoDb) FindOneAndUpdate(ctx context.Context, collection string, filter interface{},
	update, doc interface{}) error {
	if ctx == nil {
		ctx = context.TODO()
	}
	opt := options.After
	res := db.getCollection(collection).FindOneAndUpdate(ctx, filter, update,
		&options.FindOneAndUpdateOptions{ReturnDocument: &opt},
	)
	return res.Decode(doc)
}

func (db *MongoDb) UpdateMany(ctx context.Context, collection string, filter interface{},
	update interface{}) (interface{}, error) {
	if ctx == nil {
		ctx = context.TODO()
	}
	return db.getCollection(collection).UpdateMany(ctx, filter, update)
}

func (db *MongoDb) DeleteOne(ctx context.Context, collection string, filter interface{}) (interface{}, error) {
	if ctx == nil {
		ctx = context.TODO()
	}
	return db.getCollection(collection).DeleteOne(ctx, filter)
}

func (db *MongoDb) DeleteMany(ctx context.Context, collection string, filter interface{}) (interface{}, error) {
	if ctx == nil {
		ctx = context.TODO()
	}
	return db.getCollection(collection).DeleteMany(ctx, filter)
}

func (db *MongoDb) Count(ctx context.Context, collection string, filter interface{}) (int64, error) {
	if ctx == nil {
		ctx = context.TODO()
	}
	if cmp.Equal(filter, bson.M{}) {
		return db.getCollection(collection).EstimatedDocumentCount(ctx)
	}
	return db.getCollection(collection).CountDocuments(ctx, filter)
}

func (db *MongoDb) EnsureIndex(ctx context.Context, collection string, name string, keys bson.M,
	unique bool) string {
	if ctx == nil {
		ctx = context.TODO()
	}
	index := mongo.IndexModel{
		Keys:    keys,
		Options: options.Index().SetName(name).SetBackground(true).SetUnique(unique),
	}
	name, err := db.getCollection(collection).Indexes().CreateOne(ctx, index)
	if err != nil {
		fmt.Println("EnsureIndex:", collection, ":", name)
		fmt.Println(err)
	}

	return name
}

///////////////////////////////////////////////////////////////////

func (db *MongoDb) getCollection(name string) *mongo.Collection {
	return db.session.Database(db.config.Name).Collection(name)
}
