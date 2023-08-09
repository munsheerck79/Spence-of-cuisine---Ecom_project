package main

import (
	"fmt"
	"log"

	"github.com/munsheerck79/Ecom_project.git/pkg/config"
	wire "github.com/munsheerck79/Ecom_project.git/pkg/di"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Load configuration failed")
	}

	fmt.Println("jhb", config)
	server, err := wire.InitiateAPI(config)

	if err != nil {
		log.Fatal("Failed to start server")
	}
	server.Run()

}
