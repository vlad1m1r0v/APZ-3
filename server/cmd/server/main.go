package main

import (
	"github.com/vlad1m1r0v/APZ-3/server/pkg/handler"
	"github.com/vlad1m1r0v/APZ-3/server/pkg/repository"
	"github.com/vlad1m1r0v/APZ-3/server/pkg/service"
	"log"
)

func main () {
	db, err := repository.Open(repository.Config{
		Host: "localhost",
		Port: "5432",
		Username: "postgres",
		Password: "postgres",
		DBName: "virtual_machines",
		SSLMode: "disable",
	})

	if err != nil {
		log.Fatalf("failed to initialize DB: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	server := new(Server)
	if err := server.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
