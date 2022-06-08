package main

import (
	"github.com/BurntSushi/toml"
	"github.com/esabril/paimoncookies/internal/server"
	"github.com/esabril/paimoncookies/internal/service"
	"log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Panic")
		}
	}()

	var c *service.Config

	_, err := toml.DecodeFile("configs/test.toml", &c)
	if err != nil {
		log.Fatalf("Unable to read config file: %s\n", err.Error())
	}

	log.Printf("\n\nWelcome to Paimon Cookies Application. Current version: %s\n\n", c.Version)

	s := service.NewService(c)

	// maybe should change to Cobra if CLI-format will more useful
	apiEndpoint := server.GetRouter(s)

	log.Println("Running API endpoint server")

	err = apiEndpoint.Run("127.0.0.1:8087")
	if err != nil {
		log.Println("Error while starting API endpoint")
	}
}
