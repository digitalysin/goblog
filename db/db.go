package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	ORM interface {
		Set(key string, value interface{}) ORM
		Error() error
		Close() error
		Begin() ORM
		Commit() error
		Rollback() error
		Offset(offset int) ORM
		Limit(limit int) ORM
		Where(query interface{}, args ...interface{}) ORM
		Create(args interface{}) error
		Update(args interface{}) error
		Delete(model interface{}, args ...interface{}) error
	}

	mysqldb struct {
		db  *gorm.DB
		err error
	}

	MySqlOption struct {
		ConnectionString                     string
		MaxLifeTimeConnection                time.Duration
		MaxIdleConnection, MaxOpenConnection int
	}
)

func (d *mysqldb) Set(key string, value interface{}) ORM {
	var (
		db  = d.db.Set(key, value)
		err = db.Error
	)
	return &mysqldb{db, err}
}

func (d *mysqldb) Error() error {
	return d.err
}

func (d *mysqldb) Close() error {
	sql, err := d.db.DB()

	if err != nil {
		return err
	}

	if err := sql.Close(); err != nil {
		return err
	}

	return nil
}

func (d *mysqldb) Begin() ORM {
	var (
		db  = d.db.Begin()
		err = db.Error
	)
	return &mysqldb{db, err}
}

func (d *mysqldb) Commit() error {
	return d.db.Commit().Error
}

func (d *mysqldb) Rollback() error {
	return d.db.Rollback().Error
}

func (d *mysqldb) Offset(offset int) ORM {
	var (
		db  = d.db.Offset(offset)
		err = d.db.Error
	)
	return &mysqldb{db, err}
}

func (d *mysqldb) Limit(limit int) ORM {
	var (
		db  = d.db.Limit(limit)
		err = d.db.Error
	)
	return &mysqldb{db, err}
}

func (d *mysqldb) Where(query interface{}, args ...interface{}) ORM {
	var (
		db  = d.db.Where(query, args)
		err = db.Error
	)
	return &mysqldb{db, err}
}

func (d *mysqldb) Create(args interface{}) error {
	return d.db.Create(args).Error
}

func (d *mysqldb) Update(args interface{}) error {
	return d.db.Updates(args).Error

}

func (d *mysqldb) Delete(model interface{}, args ...interface{}) error {
	return d.db.Delete(model, args).Error
}

func NewMySql(option *MySqlOption) (ORM, error) {
	db, err := gorm.Open(mysql.Open(option.ConnectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sql, err := db.DB()

	if err != nil {
		return nil, err
	}

	sql.SetConnMaxLifetime(option.MaxLifeTimeConnection)
	sql.SetMaxOpenConns(option.MaxOpenConnection)
	sql.SetMaxIdleConns(option.MaxIdleConnection)

	return &mysqldb{db: db}, nil
}
