package cache

import (
	"context"
	"errors"
	"time"

	"github.com/allegro/bigcache/v3"
)

type (
	LocalMemCacheOptions struct {
		InitialMemSize  int
		MaxMemSize      int
		ExpiredInSecond int
	}

	localcache struct {
		cache *bigcache.BigCache
	}
)

func (lc *localcache) Set(_ context.Context, key string, value []byte) error {
	return lc.cache.Set(key, value)
}

func (lc *localcache) SetExp(_ context.Context, key string, value []byte, _ time.Duration) error {
	return errors.New("SetExp not implemented for local mem cache")
}

func (lc *localcache) Get(_ context.Context, key string, object interface{}) error {
	return errors.New("Get not implemented for local mem cache")
}

func (lc *localcache) GetBytes(_ context.Context, key string) ([]byte, error) {
	return lc.cache.Get(key)
}

func (lc *localcache) Keys(_ context.Context, pattern string) ([]string, error) {
	return nil, errors.New("Keys not implement for local mem cache")
}

func (lc *localcache) Ping(_ context.Context) error {
	return nil
}

func (lc *localcache) Close() error {
	return lc.cache.Close()
}

func NewLocalMemCache(opts *LocalMemCacheOptions) (Cache, error) {
	if opts.ExpiredInSecond <= 0 {
		return nil, errors.New("expired must => 0")
	}

	conf := bigcache.DefaultConfig(time.Duration(opts.ExpiredInSecond) * time.Second)
	conf.HardMaxCacheSize = opts.MaxMemSize
	conf.MaxEntrySize = opts.InitialMemSize

	bc, err := bigcache.NewBigCache(conf)

	if err != nil {
		return nil, err
	}

	return &localcache{bc}, nil
}
