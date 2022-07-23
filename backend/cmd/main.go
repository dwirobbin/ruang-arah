package main

import (
	"log"
	"ruang-arah/backend"
	"ruang-arah/backend/config"
	"ruang-arah/backend/cors"
	"ruang-arah/backend/db/conn"
	"ruang-arah/backend/db/drop"
	"ruang-arah/backend/db/migrate"
	"ruang-arah/backend/db/seed"
	"ruang-arah/backend/helper"
	"ruang-arah/backend/pkg/controller"
	"ruang-arah/backend/pkg/repository"
	"ruang-arah/backend/pkg/service"
	"ruang-arah/backend/router"
)

func main() {
	db, err := conn.Connect()
	helper.PanicIfError(err)
	defer db.Close()

	drop.Drop(db)
	migrate.Migrate(db)
	seed.Seed(db)

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	controllers := controller.NewController(services)

	router := router.NewRouter(controllers)

	srv := new(backend.Server)
	if err := srv.Run(config.API_PORT, cors.AllowOrigin(router)); err != nil {
		log.Fatal("Error occured while running server: ", err.Error())
	}
}
