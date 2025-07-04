package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type (
	Cache interface {
		Set(ctx context.Context, key string, value []byte) error
		SetExp(ctx context.Context, key string, value []byte, exp time.Duration) error
		Get(ctx context.Context, key string, object interface{}) error
		GetBytes(ctx context.Context, key string) ([]byte, error)
		Incr(ctx context.Context, key string) error
		Decr(ctx context.Context, key string) error
		Del(ctx context.Context, key string) error
		Keys(ctx context.Context, pattern string) ([]string, error)
		Ping(ctx context.Context) error
		Close() error
	}

	Option struct {
		Address, UserName, Password                        string
		DB, PoolSize, MinIdleConn                          int
		DialTimeout, ReadTimeout, WriteTimeout, MaxConnAge time.Duration
	}

	cch struct {
		cache *redis.Client
	}
)

func (c *cch) Set(ctx context.Context, key string, value []byte) error {
	return c.SetExp(ctx, key, value, 0)
}

func (c *cch) SetExp(ctx context.Context, key string, value []byte, exp time.Duration) error {
	var (
		status = c.cache.Set(ctx, key, value, exp)
	)
	return status.Err()
}

func (c *cch) Get(ctx context.Context, key string, object interface{}) error {
	var (
		status = c.cache.Get(ctx, key)
	)

	if err := status.Err(); err != nil {
		return err
	}

	return status.Scan(object)
}

func (c *cch) GetBytes(ctx context.Context, key string) ([]byte, error) {
	var (
		status = c.cache.Get(ctx, key)
	)

	if err := status.Err(); err != nil {
		return nil, err
	}

	return status.Bytes()
}

func (c *cch) Incr(ctx context.Context, key string) error {
	return c.cache.Incr(ctx, key).Err()
}

func (c *cch) Decr(ctx context.Context, key string) error {
	return c.cache.Decr(ctx, key).Err()
}

func (c *cch) Del(ctx context.Context, key string) error {
	return c.cache.Del(ctx, key).Err()
}

func (c *cch) Keys(ctx context.Context, pattern string) ([]string, error) {
	var (
		res = c.cache.Keys(ctx, pattern)
	)

	return res.Result()
}

func (c *cch) Ping(ctx context.Context) error {
	return c.cache.Ping(ctx).Err()
}

func (c *cch) Close() error {
	return c.cache.Close()
}

func New(option *Option) (Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         option.Address,
		Username:     option.UserName,
		Password:     option.Password,
		DB:           option.DB,
		DialTimeout:  option.DialTimeout,
		ReadTimeout:  option.ReadTimeout,
		WriteTimeout: option.WriteTimeout,
		MaxConnAge:   option.MaxConnAge,
		PoolSize:     option.PoolSize,
		MinIdleConns: option.MinIdleConn,
	})

	return &cch{client}, nil
}
