package person

import (
	"context"
	"fmt"
	"my-rest-api/utils/error_utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	PersonRepo personRepoInterface = &personRepo{}
)

const (
	dbName         = "personsdb"
	collectionName = "person"
)

type personRepoInterface interface {
	Initialize(*mongo.Client) *mongo.Client
	Get(bson.M) ([]bson.M, error_utils.MessageErr)
	GetAll() ([]bson.M, error_utils.MessageErr)
	Create(*Person) (*mongo.InsertOneResult, error_utils.MessageErr)
	Update(primitive.M, bson.M) (*mongo.UpdateResult, error_utils.MessageErr)
	Delete(bson.M) (*mongo.DeleteResult, error_utils.MessageErr)
}

type personRepo struct {
	db *mongo.Client
}

func (pr *personRepo) Initialize(client *mongo.Client) *mongo.Client {
	pr.db = client
	return pr.db
}

func NewPersonRepository(db *mongo.Client) personRepoInterface {
	return &personRepo{db: db}
}

func (pr *personRepo) Get(personId bson.M) ([]bson.M, error_utils.MessageErr) {
	errDB := pr.db.Ping(context.Background(), readpref.Primary())
	if errDB != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("Error when trying to connect to database: %s", errDB.Error()))
	}

	collection := pr.db.Database(dbName).Collection(collectionName)
	cur, err := collection.Find(context.Background(), personId)
	defer cur.Close(context.Background())

	if err != nil {
		return nil, error_utils.NewInternalServerError("Error getting data from database")
	}

	var results []bson.M

	cur.All(context.Background(), &results)

	if results == nil {
		return nil, error_utils.NewNotFoundError(fmt.Sprintf("Not found results"))
	}

	return results, nil
}

func (pr *personRepo) GetAll() ([]bson.M, error_utils.MessageErr) {
	errDB := pr.db.Ping(context.Background(), readpref.Primary())
	if errDB != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("Error when trying to connect to database: %s", errDB.Error()))
	}

	collection := pr.db.Database(dbName).Collection(collectionName)
	var results []bson.M
	cur, err := collection.Find(context.Background(), bson.M{})
	defer cur.Close(context.Background())

	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("Error getting collection from database %s", err.Error()))
	}

	cur.All(context.Background(), &results)

	if results == nil {
		return nil, error_utils.NewNotFoundError(fmt.Sprintf("Not found results"))
	}

	return results, nil
}

func (pr *personRepo) Create(person *Person) (*mongo.InsertOneResult, error_utils.MessageErr) {
	errDB := pr.db.Ping(context.Background(), readpref.Primary())
	if errDB != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("Error when trying to connect to database: %s", errDB.Error()))
	}

	collection := pr.db.Database(dbName).Collection(collectionName)
	res, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("Error when trying to update: %s", err.Error()))
	}

	return res, nil
}

func (pr *personRepo) Update(ppId primitive.M, ppObj bson.M) (*mongo.UpdateResult, error_utils.MessageErr) {
	errDB := pr.db.Ping(context.Background(), readpref.Primary())
	if errDB != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("Error when trying to connect to database: %s", errDB.Error()))
	}

	collection := pr.db.Database(dbName).Collection(collectionName)
	person, err := collection.UpdateOne(context.Background(), ppId, ppObj)

	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("Error processing the update %s", err))
	}

	return person, nil
}

func (pr *personRepo) Delete(personId bson.M) (*mongo.DeleteResult, error_utils.MessageErr) {
	errDB := pr.db.Ping(context.Background(), readpref.Primary())
	if errDB != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("Error when trying to connect to database: %s", errDB.Error()))
	}

	collection := pr.db.Database(dbName).Collection(collectionName)
	res, err := collection.DeleteOne(context.Background(), personId)
	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("Error processing the delete %s", err))
	}
	return res, nil
}
