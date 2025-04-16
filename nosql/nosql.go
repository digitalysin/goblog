package nosql

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	Entity interface {
		CollectionName() string
	}

	Repository[E Entity, T ~[]E] interface {
		FindOne(ctx context.Context, filter interface{}) (E, error)
		FindMany(ctx context.Context, filter interface{}) (T, error)
	}

	MongoRepository[E Entity, T ~[]E] interface {
		Repository[E, T]
		GetClient() *mongo.Client
		GetDatabase() string
	}

	AbstractMongoCrudRepository[E Entity, T ~[]E] struct {
		Client   *mongo.Client
		Database string
	}
)

func (impl *AbstractMongoCrudRepository[E, T]) GetClient() *mongo.Client {
	return impl.Client
}

func (impl *AbstractMongoCrudRepository[E, T]) GetDatabase() string {
	return impl.Database
}

func (impl *AbstractMongoCrudRepository[E, T]) FindOne(ctx context.Context, filter interface{}) (E, error) {
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

func (impl *AbstractMongoCrudRepository[E, T]) FindMany(ctx context.Context, filter interface{}) (T, error) {
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
