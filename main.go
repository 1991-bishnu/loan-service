package main

import (
	"fmt"
	"log"

	"github.com/1991-bishnu/loan-service/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(config)

}
