package main

import (
	"order-service/adapter"
	"order-service/config"
	"order-service/controllers"
	"order-service/middleware"
	shopRepo "order-service/repositories"
	"order-service/routes"
	"order-service/server"
	orderService "order-service/services"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	config.Load()
	adapter.LoadMultipleDatabase()
	adapter.LoadMigrationsDatabase()
}

func main() {
	listDb := viper.GetString("orderservice.db.name") // remove this if you won't use database
	dbName := strings.Split(listDb, ",")              // remove this if you won't use database
	db := adapter.DBConnection(dbName[0])             // remove this if you won't use database

	log := middleware.Logrus()
	timezone := middleware.Timezone()

	//list repo will be here
	orderRepo := shopRepo.NewOrderRepository(db)

	//list service will be here
	orderService := orderService.NewOrderService(orderRepo, db)

	//list controller will be here
	controller := controllers.NewController(orderService)
	mdl := middleware.NewMiddleware(log, timezone)

	router := routes.NewRouter(mdl, controller)

	server.Wrapper(router.Route())

}
