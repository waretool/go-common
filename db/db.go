package db

import (
	"fmt"
	"github.com/waretool/go-common/env"
	"github.com/waretool/go-common/logger"
	"gorm.io/gorm/schema"
	"time"

	"github.com/heptiolabs/healthcheck"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type Database interface {
	Conn() *gorm.DB
	HealthCheck() healthcheck.Check
}

type database struct {
	connection *gorm.DB
	user       string
	password   string
	host       string
	port       string
	schema     string
}

func (db *database) Conn() *gorm.DB {
	return db.connection
}

func (db *database) HealthCheck() healthcheck.Check {
	sqlDb, err := db.connection.DB()
	if err != nil {
		logger.Errorf("unable to create a healthcheck for database connection.")
	}
	return healthcheck.DatabasePingCheck(sqlDb, time.Second)
}

func createDatabase() Database {
	var db *gorm.DB = nil

	retry := env.GetEnv("DB_CONN_RETRY", 10)
	dbUser := env.GetEnv("DB_USER", "db-user")
	dbPassword := env.GetEnv("DB_PASSWORD", "db-password")
	dbHost := env.GetEnv("DB_HOST", "db-host")
	dbPort := env.GetEnv("DB_PORT", "db-port")
	dbSchema := env.GetEnv("DB_DATABASE", "db-schema")

	var connected = false
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbSchema)
	for i := 0; !connected && i < retry; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Silent),
			NamingStrategy: schema.NamingStrategy{
				NoLowerCase: true,
			},
		})
		if err != nil {
			logger.Errorf("cannot connect to db '%s'. attempt number %d of %d failed.", dbSchema, i+1, retry)
			time.Sleep(time.Duration(i*i) * time.Second)
			continue
		}
		connected = true
	}

	if !connected {
		logger.Error("cannot connect to database. exiting...")
		panic(1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatalf("connection pool cannot be set up due to: %s", err)
	}

	maxConnection := env.GetEnv("MYSQL_MAX_OPEN_CONN", 150)
	sqlDB.SetMaxOpenConns(maxConnection)
	sqlDB.SetMaxIdleConns(15)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &database{
		connection: db,
	}
}
