package main

import (
	"log"

	"github.com/1991-bishnu/loan-service/config"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

}
