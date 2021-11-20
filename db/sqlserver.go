package db

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
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
