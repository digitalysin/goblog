package db

import (
	"context"
)

type (
	ORM interface {
		Set(key string, value interface{}) ORM
		Error() error
		Close() error
		Begin() ORM
		Commit() error
		Rollback() error
		Offset(offset int64) ORM
		Limit(limit int64) ORM
		First(object interface{}) error
		Last(object interface{}) error
		Find(object interface{}) error
		Model(value interface{}) ORM
		Select(query interface{}, args ...interface{}) ORM
		OmitAssoc() ORM
		Where(query interface{}, args ...interface{}) ORM
		Order(value interface{}) ORM
		Create(args interface{}) error
		Update(args interface{}) error
		UpdateColumns(args interface{}) error
		Delete(model interface{}, args ...interface{}) error
		WithContext(ctx context.Context) ORM
		Raw(query string, args ...interface{}) ORM
		Exec(query string, args ...interface{}) ORM
		Scan(object interface{}) error
		Preload(assoc string) ORM
		Joins(assoc string) ORM
		Ping() error
	}
)
