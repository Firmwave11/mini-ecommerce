package main

import (
	"strings"
	"user-service/adapter"
	"user-service/config"
	"user-service/controllers"
	"user-service/middleware"
	tokenRepo "user-service/repositories/token"
	userRepo "user-service/repositories/user"
	"user-service/routes"
	tokenService "user-service/services/token"
	userService "user-service/services/user"
	"user-service/utils/server"

	"github.com/spf13/viper"
)

func init() {
	config.Load()
	adapter.LoadMultipleDatabase()
	adapter.LoadMigrationsDatabase()
}

func main() {
	listDb := viper.GetString("userservice.db.name") // remove this if you won't use database
	dbName := strings.Split(listDb, ",")             // remove this if you won't use database
	db := adapter.DBConnection(dbName[0])            // remove this if you won't use database

	log := middleware.Logrus()
	timezone := middleware.Timezone()

	//list repo will be here
	tokenRepo := tokenRepo.NewTokenRepository(db)
	userRepo := userRepo.NewUserRepository(db)

	//list service will be here
	tokenService := tokenService.NewTokenService(tokenRepo, db)
	userService := userService.NewUserService(tokenRepo, userRepo, db)

	//list controller will be here
	controller := controllers.NewController(userService, tokenService)
	mdl := middleware.NewMiddleware(log, timezone)

	router := routes.NewRouter(mdl, controller)

	server.Wrapper(router.Route())

}
