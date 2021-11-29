package server

import "log"

func main () {
	server := new(Server)
	if err := server.Run("8000"); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
