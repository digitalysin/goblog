package db

import (
	glog "github.com/digitalysin/goblog/logger"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type (
	SqlServerOption struct {
		ConnectionString                     string
		MaxLifeTimeConnection                time.Duration
		MaxIdleConnection, MaxOpenConnection int
		Logger                               glog.Logger
	}
)

func NewSqlServer(option *SqlServerOption) (ORM, error) {
	var (
		opts = &gorm.Config{}
	)

	if option.Logger != nil {
		opts.Logger = logger.New(option.Logger, logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			Colorful:                  false,
			IgnoreRecordNotFoundError: false,
		})
	}

	db, err := gorm.Open(sqlserver.Open(option.ConnectionString), opts)

	if err != nil {
		return nil, err
	}

	sql, err := db.DB()

	sql.SetConnMaxLifetime(option.MaxLifeTimeConnection)
	sql.SetMaxOpenConns(option.MaxOpenConnection)
	sql.SetMaxIdleConns(option.MaxIdleConnection)

	return &mysqldb{db: db}, nil
}
