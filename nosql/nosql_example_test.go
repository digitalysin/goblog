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
		repo    = NewMongoRepository(client, "tix_flight_search_omg")
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

	updated, err := repo.Update(ctx, airline, bson.D{{
		"$set", bson.D{{"code", "new updated code"}, {"name", "new updated name"}},
	}})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(updated)
}

type (
	Airline struct {
		ID   bson.ObjectID `bson:"_id"`
		Code string        `bson:"code"`
		Name string        `bson:"name"`
	}
	Airlines []*Airline
)

func (impl *Airline) CollectionName() string {
	return "airline"
}

func (impl *Airline) GetID() bson.ObjectID { return impl.ID }

func NewMongoRepository(client *mongo.Client, database string) *AbstractMongoCrudRepository[bson.ObjectID, *Airline, Airlines] {
	return &AbstractMongoCrudRepository[bson.ObjectID, *Airline, Airlines]{
		Client:   client,
		Database: database,
	}
}
