package nosql

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	AbstractMongoCrudRepository[ID any, E Entity[ID], T ~[]E] struct {
		Client   *mongo.Client
		Database string
	}
)

func (impl *AbstractMongoCrudRepository[ID, E, T]) GetClient() *mongo.Client {
	return impl.Client
}

func (impl *AbstractMongoCrudRepository[ID, E, T]) GetDatabase() string {
	return impl.Database
}

func (impl *AbstractMongoCrudRepository[ID, E, T]) FindByID(ctx context.Context, id ID) (E, error) {
	e, err := impl.FindOne(ctx, bson.D{{"_id", id}})

	if err != nil {
		return e, errors.Wrapf(err, "failed to find by id %s", id)
	}

	return e, nil
}

func (impl *AbstractMongoCrudRepository[ID, E, T]) FindOne(ctx context.Context, filter interface{}) (E, error) {
	var (
		entity     E
		client     = impl.GetClient()
		collection = client.Database(impl.GetDatabase()).Collection(entity.CollectionName())
		err        = collection.FindOne(ctx, filter).Decode(&entity)
	)

	if err != nil {
		return entity, errors.Wrapf(err, "failed to find one in collection %s", entity.CollectionName())
	}

	return entity, nil
}

func (impl *AbstractMongoCrudRepository[ID, E, T]) FindMany(ctx context.Context, filter interface{}) (T, error) {
	var (
		entity     E
		entities   = make([]E, 0)
		client     = impl.GetClient()
		collection = client.Database(impl.GetDatabase()).Collection(entity.CollectionName())
		err        = collection.FindOne(ctx, filter).Decode(&entity)
	)

	if err != nil {
		return entities, errors.Wrapf(err, "failed to find many in collection %s", entity.CollectionName())
	}

	return entities, nil
}

func (impl *AbstractMongoCrudRepository[ID, E, T]) Create(ctx context.Context, entity E) (E, error) {
	var (
		client     = impl.GetClient()
		collection = client.Database(impl.GetDatabase()).Collection(entity.CollectionName())
		_, err     = collection.InsertOne(ctx, entity)
	)

	if err != nil {
		return entity, errors.Wrapf(err, "failed to create in collection %s", entity.CollectionName())
	}

	return entity, nil
}

func (impl *AbstractMongoCrudRepository[ID, E, T]) Update(ctx context.Context, entity E, update interface{}) error {
	var (
		client     = impl.GetClient()
		collection = client.Database(impl.GetDatabase()).Collection(entity.CollectionName())
		_, err     = collection.UpdateOne(ctx, bson.D{{"_id", entity.GetID()}}, update)
	)

	if err != nil {
		return errors.Wrapf(err, "failed to update one in collection %s", entity.CollectionName())
	}

	return nil
}

func (impl *AbstractMongoCrudRepository[ID, E, T]) Delete(ctx context.Context, entity E) error {
	var (
		client     = impl.GetClient()
		collection = client.Database(impl.GetDatabase()).Collection(entity.CollectionName())
		_, err     = collection.DeleteOne(ctx, bson.D{{"_id", entity.GetID()}})
	)

	if err != nil {
		return errors.Wrapf(err, "failed to delete in collection %s", entity.CollectionName())
	}

	return nil
}

func (impl *AbstractMongoCrudRepository[ID, E, T]) DeleteMany(ctx context.Context, filter interface{}) error {
	var (
		entity     E
		client     = impl.GetClient()
		collection = client.Database(impl.GetDatabase()).Collection(entity.CollectionName())
		_, err     = collection.DeleteMany(ctx, filter)
	)

	if err != nil {
		return errors.Wrapf(err, "failed to delete many in collection %s", entity.CollectionName())
	}

	return nil
}
