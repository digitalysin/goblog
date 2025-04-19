package nosql

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"testing"
)

func TestNoSQL(t *testing.T) {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Panic(err)
	}

	var (
		ctx = context.Background()
	)

	defer client.Disconnect(ctx)

	var (
		repo, _ = NewMongoRepository(client, "tix_flight_search_omg")
		airline = &Airline{
			ID:   bson.NewObjectID(),
			Code: "some code from me",
			Name: "hello world",
		}
	)

	created, err := repo.Create(ctx, airline)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(created)

	airline, err = repo.FindOne(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(airline)

	airline = &Airline{
		ID:   airline.ID,
		Code: "the new one",
		Name: "hello world the new airline",
	}

	err = repo.Update(ctx, airline, bson.D{{
		"$set", bson.D{{"code", "new updated code"}, {"name", "new updated name"}},
	}})

	if err != nil {
		log.Fatal(err)
	}

}

type (
	Airline struct {
		ID   bson.ObjectID `bson:"_id"`
		Code string        `bson:"code"`
		Name string        `bson:"name"`
	}
	Airlines          []*Airline
	AirlineRepository interface {
		MongoRepository[bson.ObjectID, *Airline, Airlines]
	}

	airlineRepositoryImpl struct {
		AbstractMongoCrudRepository[bson.ObjectID, *Airline, Airlines]
	}
)

func (impl *Airline) CollectionName() string { return "airline" }
func (impl *Airline) GetID() bson.ObjectID   { return impl.ID }

func NewMongoRepository(client *mongo.Client, database string) (AirlineRepository, error) {
	return &airlineRepositoryImpl{
		AbstractMongoCrudRepository: AbstractMongoCrudRepository[bson.ObjectID, *Airline, Airlines]{
			Client:   client,
			Database: database,
		},
	}, nil
}
