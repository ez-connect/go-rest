package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
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

func (db *MongoDb) GetCursor(collection string,
	filter, sort interface{},
	skip, limit int64,
	projection interface{}) (interface{}, error) {

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

	return db.getCollection(collection).Find(context.TODO(), filter, opts)
}

func (db *MongoDb) Find(collection string,
	filter interface{}, option FindOption,
	projection, docs interface{}) error {

	var sort bson.M
	if option.Sort != "" {
		if option.Order == "desc" {
			sort = bson.M{option.Sort: -1}
		} else {
			sort = bson.M{option.Sort: 1}
		}
	}

	res, err := db.GetCursor(
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
	defer cur.Close(context.TODO())
	return cur.All(context.TODO(), docs)
}

func (db *MongoDb) Distinct(collection, fieldName string,
	filter interface{}) ([]string, error) {
	res, err := db.getCollection(collection).Distinct(context.TODO(), fieldName, filter)
	values := make([]string, len(res))
	for i := range res {
		values[i] = res[i].(string)
	}
	return values, err
}

func (db *MongoDb) Aggregate(collection string, pipeline interface{},
	docs interface{}) error {
	cur, err := db.getCollection(collection).Aggregate(context.TODO(), pipeline)
	if err == nil {
		defer cur.Close(context.TODO())
		return cur.All(context.TODO(), docs)
	}
	return err
}

func (db *MongoDb) FindOne(collection string, filter interface{},
	projection interface{}, doc interface{}) error {
	opts := options.FindOne()
	if projection != nil {
		opts.SetProjection(projection)
	}
	res := db.getCollection(collection).FindOne(context.TODO(), filter, opts)
	return res.Decode(doc)
}

func (db *MongoDb) AggregateOne(collection string, pipeline interface{},
	doc interface{}) error {
	cur, err := db.getCollection(collection).Aggregate(context.TODO(), pipeline)
	if err != nil {
		return err
	}

	defer cur.Close(context.TODO())
	if cur.Next(context.TODO()) {
		return cur.Decode(doc)
	}

	return nil
}

func (db *MongoDb) Insert(collection string, doc interface{}) (interface{}, error) {
	return db.getCollection(collection).InsertOne(context.TODO(), doc)
}

func (db *MongoDb) InsertMany(collection string, docs []interface{}) ([]interface{}, error) {
	res, err := db.getCollection(collection).InsertMany(context.TODO(), docs)
	return res.InsertedIDs, err
}

func (db *MongoDb) UpdateOne(collection string, filter interface{},
	doc interface{}) (interface{}, error) {
	return db.getCollection(collection).UpdateOne(context.TODO(), filter, bson.M{"$set": doc})
}

func (db *MongoDb) DeleteOne(collection string, filter interface{}) (interface{}, error) {
	return db.getCollection(collection).DeleteOne(context.TODO(), filter)
}

func (db *MongoDb) DeleteMany(collection string, filter interface{}) (interface{}, error) {
	return db.getCollection(collection).DeleteMany(context.TODO(), filter)
}

func (db *MongoDb) Count(collection string, filter interface{}) (int64, error) {
	if cmp.Equal(filter, bson.M{}) {
		return db.getCollection(collection).EstimatedDocumentCount(context.TODO())
	}
	return db.getCollection(collection).CountDocuments(context.TODO(), filter)
}

func (db *MongoDb) EnsureIndex(collection string, name string, keys bson.M,
	unique bool) string {
	index := mongo.IndexModel{
		Keys:    keys,
		Options: options.Index().SetName(name).SetBackground(true).SetUnique(unique),
	}
	name, err := db.getCollection(collection).Indexes().CreateOne(context.TODO(), index)
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
