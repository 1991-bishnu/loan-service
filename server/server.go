package server

import (
	"fmt"

	"github.com/1991-bishnu/loan-service/config"
)

func Start(conf *config.AppConfig) error {
	router := NewRouter(conf)
	err := router.Run(conf.Server.Address)
	if err != nil {
		return fmt.Errorf("failed to run server error: %w", err)
	}
	return nil
}
