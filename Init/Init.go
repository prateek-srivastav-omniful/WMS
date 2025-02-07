package Init

import (
	"context"
	"log"

	"github.com/prateek-srivastav-omniful/wms-service/controllers"
	"github.com/prateek-srivastav-omniful/wms-service/database"
	"github.com/prateek-srivastav-omniful/wms-service/models"
	"gorm.io/gorm"
	// "github.com/omniful/go_commons/"
)

func InitializeDb(ctx context.Context) {
	connector := &database.Connection{}
	connector.DB = connector.ConnectPostgres(ctx)
	// Initialze Tables
	// InitializeTables(connector.DB)
	// pass this db instance where needed
	controllers.SetDbInstance(connector.DB)

}
func InitializeTables(db *gorm.DB) {
	err := db.AutoMigrate(&models.Hub{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
