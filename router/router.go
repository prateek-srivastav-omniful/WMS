package router

import (
	"context"

	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/log"
	"github.com/prateek-srivastav-omniful/wms-service/controllers"
)

func PublicRoutes(ctx context.Context, server *http.Server) error {

	apiV1 := server.Engine.Group("/api/v1", log.RequestLogMiddleware(log.MiddlewareOptions{
		Format:      config.GetString(ctx, "log.format"),
		Level:       config.GetString(ctx, "log.level"),
		LogRequest:  config.GetBool(ctx, "log.request"),
		LogResponse: config.GetBool(ctx, "log.response"),
	}))
	apiV2 := server.Engine.Group("/api/v2", log.RequestLogMiddleware(log.MiddlewareOptions{
		Format:      config.GetString(ctx, "log.format"),
		Level:       config.GetString(ctx, "log.level"),
		LogRequest:  config.GetBool(ctx, "log.request"),
		LogResponse: config.GetBool(ctx, "log.response"),
	}))
	HubV1 := apiV1.Group("/api/v1/hub")

	HubV1.POST("", controllers.CreateHub)
	HubV1.GET("", controllers.GetHubs)

	inventory := apiV1.Group("/inventory")
	inventory.POST("", controllers.CreateInventory)
	inventory.GET("/:id", controllers.GetInventoryByID)
	inventory.GET("", controllers.GetAllInventory)
	inventory.PUT("/:id", controllers.UpdateInventory)
	inventory.DELETE("/:id", controllers.DeleteInventory)

	// Tenant Routes
	tenant := apiV1.Group("/tenant")
	tenant.POST("", controllers.CreateTenant)
	tenant.GET("/:id", controllers.GetTenantByID)
	tenant.GET("", controllers.GetAllTenants)
	tenant.PUT("/:id", controllers.UpdateTenant)
	tenant.DELETE("/:id", controllers.DeleteTenant)

	sku := apiV1.Group("/sku")
	sku.POST("", controllers.CreateSKU)
	sku.GET("/:id", controllers.GetSKUByID)
	sku.GET("", controllers.GetAllSKUs)
	sku.PUT("/:id", controllers.UpdateSKU)
	sku.DELETE("/:id", controllers.DeleteSKU)

	validateHub := apiV2.Group("/validate-hub")
	validateHub.GET("", controllers.ValidateHub)
	validateSku := apiV2.Group("/validate-sku")
	validateSku.GET("", controllers.ValidateSKU)

	return nil
}
