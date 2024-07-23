package main

import (
	"shop-service/adapter"
	"shop-service/config"
	"shop-service/controllers"
	"shop-service/middleware"
	shopRepo "shop-service/repositories"
	"shop-service/routes"
	"shop-service/server"
	shopService "shop-service/services"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	config.Load()
	adapter.LoadMultipleDatabase()
	adapter.LoadMigrationsDatabase()
}

func main() {
	listDb := viper.GetString("shopservice.db.name") // remove this if you won't use database
	dbName := strings.Split(listDb, ",")             // remove this if you won't use database
	db := adapter.DBConnection(dbName[0])            // remove this if you won't use database

	log := middleware.Logrus()
	timezone := middleware.Timezone()

	//list repo will be here
	shopRepo := shopRepo.NewShopRepository(db)

	//list service will be here
	shopService := shopService.NewShopService(shopRepo, db)

	//list controller will be here
	controller := controllers.NewController(shopService)
	mdl := middleware.NewMiddleware(log, timezone)

	router := routes.NewRouter(mdl, controller)

	server.Wrapper(router.Route())

}
