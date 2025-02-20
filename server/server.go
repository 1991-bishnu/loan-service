package server

import "github.com/1991-bishnu/loan-service/config"

func Start(conf *config.AppConfig) {
	r := NewRouter(conf)
	r.Run(conf.Server.Address)
}
