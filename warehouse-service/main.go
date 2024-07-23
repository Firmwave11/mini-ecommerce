package main

import (
	"strings"
	"warehouse-service/adapter"
	"warehouse-service/config"
	"warehouse-service/controllers"
	"warehouse-service/middleware"
	warehouseRepo "warehouse-service/repositories"
	"warehouse-service/routes"
	"warehouse-service/server"
	warehouseService "warehouse-service/services"

	"github.com/spf13/viper"
)

func init() {
	config.Load()
	adapter.LoadMultipleDatabase()
	adapter.LoadMigrationsDatabase()
}

func main() {
	listDb := viper.GetString("warehouseservice.db.name") // remove this if you won't use database
	dbName := strings.Split(listDb, ",")                  // remove this if you won't use database
	db := adapter.DBConnection(dbName[0])                 // remove this if you won't use database

	log := middleware.Logrus()
	timezone := middleware.Timezone()

	//list repo will be here
	warehouseRepo := warehouseRepo.NewWarehouseRepository(db)

	//list service will be here
	warehouseService := warehouseService.NewWarehouseService(warehouseRepo, db)

	//list controller will be here
	controller := controllers.NewController(warehouseService)
	mdl := middleware.NewMiddleware(log, timezone)

	router := routes.NewRouter(mdl, controller)

	server.Wrapper(router.Route())

}
