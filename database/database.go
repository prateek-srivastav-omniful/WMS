package database

import (
	"context"
	"time"

	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/db/sql/postgres"
	"gorm.io/gorm"
)

type PostgresConnection interface {
	ConnectMongo(context.Context)
}

type Connection struct {
	// DB *mongo.Client
	DB *gorm.DB
}

func (C *Connection) ConnectPostgres(con context.Context) *gorm.DB {
	connIdleTimeout := 10 * time.Minute
	debugMode := config.GetBool(con, "postgresql.debugMode")

	// Master config i.e. - Write endpoint
	masterConfig := postgres.DBConfig{
		Host:               "localhost",
		Port:               "5432",
		Username:           "postgres",
		Password:           "prateek987",
		Dbname:             "postgres",
		MaxOpenConnections: 10,
		MaxIdleConnections: 2,
		ConnMaxLifetime:    connIdleTimeout,
		DebugMode:          debugMode,
	}

	// Slave config i.e. - array with read endpoints
	slavesConfig := make([]postgres.DBConfig, 0)
	DbCluster := postgres.InitializeDBInstance(masterConfig, &slavesConfig)
	MasterDb := DbCluster.GetMasterDB(con)
	return MasterDb
}
