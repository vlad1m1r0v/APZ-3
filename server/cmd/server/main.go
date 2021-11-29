package main

import (
	"github.com/vlad1m1r0v/APZ-3/server/pkg/handler"
	"log"
)

func main () {
	handler := new(handler.Handler)
	server := new(Server)
	if err := server.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
