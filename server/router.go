package server

import (
	"github.com/1991-bishnu/loan-service/config"
	"github.com/1991-bishnu/loan-service/controller"
	"github.com/1991-bishnu/loan-service/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(conf *config.AppConfig) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controller.HealthController)

	router.GET("", health.Welcome)
	router.GET("/health", health.Status)

	router.Use(middleware.AuthMiddleware(conf))

	return router

}
