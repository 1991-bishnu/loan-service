package main

import (
	"log"

	"github.com/1991-bishnu/loan-service/config"
	"github.com/1991-bishnu/loan-service/db"
	"github.com/1991-bishnu/loan-service/db/seed"
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

	// Optional TODO: Move to cmd
	err = db.MigrateDB()
	if err != nil {
		log.Fatal(err)
	}

	// Optional TODO: Move to cmd
	err = seed.Seed(db.GetDB())
	if err != nil {
		log.Print(err)
	}

	err = server.Start(conf)
	if err != nil {
		log.Fatal(err)
	}
}
