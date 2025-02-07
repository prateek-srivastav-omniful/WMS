package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/omniful/go_commons/db/sql/migration"
	"github.com/omniful/go_commons/http"
	"github.com/prateek-srivastav-omniful/wms-service/Init"
	"github.com/prateek-srivastav-omniful/wms-service/router"
)

func initialize(ctx context.Context) {

	// Initialize Database
	Init.InitializeDb(ctx)
	// Init.InitializeSqs(ctx)
	// Initialize Routers

}

const (
	modeWorker     = "worker"
	modeHttp       = "http"
	modeMigration  = "migration"
	upMigration    = "up"
	downMigration  = "down"
	forceMigration = "force"
)

func main() {
	// Create background context
	ctx := context.Background()
	go initialize(ctx)

	runMigration(ctx, "up", "5")

	err := runserver(ctx)
	if err != nil {
		fmt.Println(err)
	}

}
func runserver(ctx context.Context) error {
	server := http.InitializeServer(":8582", 10*time.Second, 10*time.Second, 77*time.Second)
	fmt.Println("Server Initialized")

	//routes
	err := router.PublicRoutes(ctx, server)
	if err != nil {
		return err
	}

	err = server.StartServer("WMS")
	if err != nil {
		fmt.Println("Error Starting Server:", err)
		return err
	}
	return nil
}

func runMigration(ctx context.Context, migrationType string, number string) {

	m, err := migration.InitializeMigrate("file://deployment/migration", "postgres://postgres:prateek987@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// migrationType = "force"
	switch migrationType {
	case upMigration:
		err = m.Up()
		if err != nil {
			panic(err)
		}
		break
	case downMigration:
		err = m.Down()
		if err != nil {
			panic(err)
		}
		break
	case forceMigration:
		version, parseErr := strconv.Atoi(number)
		if parseErr != nil {
			panic(parseErr)
		}

		err = m.ForceVersion(version)
		if err != nil {
			return
		}
		// break
	default:
		err = m.Up()
		if err != nil {
			panic(err)
		}
		break
	}
}
