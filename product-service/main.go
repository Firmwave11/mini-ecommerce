package main

import (
	"product-service/adapter"
	"product-service/config"
	"product-service/controllers"
	"product-service/middleware"
	productRepo "product-service/repositories"
	"product-service/routes"
	"product-service/server"
	productService "product-service/services"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	config.Load()
	adapter.LoadMultipleDatabase()
	adapter.LoadMigrationsDatabase()
}

func main() {
	listDb := viper.GetString("productservice.db.name") // remove this if you won't use database
	dbName := strings.Split(listDb, ",")                // remove this if you won't use database
	db := adapter.DBConnection(dbName[0])               // remove this if you won't use database

	log := middleware.Logrus()
	timezone := middleware.Timezone()

	//list repo will be here
	productRepo := productRepo.NewProductRepository(db)

	//list service will be here
	productService := productService.NewProductService(productRepo, db)

	//list controller will be here
	controller := controllers.NewController(productService)
	mdl := middleware.NewMiddleware(log, timezone)

	router := routes.NewRouter(mdl, controller)

	server.Wrapper(router.Route())

}
