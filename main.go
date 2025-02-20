package main

import (
	"log"

	"github.com/1991-bishnu/loan-service/config"
	"github.com/1991-bishnu/loan-service/db"
	"github.com/1991-bishnu/loan-service/server"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Init(conf)
	if err != nil {
		log.Fatal(err)
	}

	err = db.MigrateDB()
	if err != nil {
		log.Fatal(err)
	}

	err = server.Start(conf)
	if err != nil {
		log.Fatal(err)
	}

}
