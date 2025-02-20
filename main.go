package main

import (
	"log"

	"github.com/1991-bishnu/loan-service/config"
	"github.com/1991-bishnu/loan-service/server"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	server.Start(config)
}
