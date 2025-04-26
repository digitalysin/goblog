package nosql

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	Repository[ID any, E Entity[ID], T ~[]E] interface {
		FindByID(ctx context.Context, id ID) (E, error)
		FindOne(ctx context.Context, filter interface{}) (E, error)
		FindMany(ctx context.Context, filter interface{}) (T, error)
		Create(ctx context.Context, entity E) (E, error)
		Update(ctx context.Context, e E, update interface{}) error
		Delete(ctx context.Context, e E) error
		DeleteMany(ctx context.Context, filter interface{}) error
	}

	MongoRepository[ID any, E Entity[ID], T ~[]E] interface {
		Repository[ID, E, T]
		GetClient() *mongo.Client
		GetDatabase() string
	}
)
