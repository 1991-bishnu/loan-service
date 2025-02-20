package server

import (
	"github.com/1991-bishnu/loan-service/config"
	"github.com/1991-bishnu/loan-service/controller"
	"github.com/1991-bishnu/loan-service/db"
	"github.com/1991-bishnu/loan-service/middleware"
	"github.com/1991-bishnu/loan-service/service"
	"github.com/1991-bishnu/loan-service/store"
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

	userStoreObj := store.NewUser(db.GetDB())
	loanStoreObj := store.NewLoan(db.GetDB())
	employeeStoreObj := store.NewEmployee(db.GetDB())
	documentStoreObj := store.NewDocument(db.GetDB())
	investorStoreObj := store.NewInvestor(db.GetDB())
	investmentStoreObj := store.NewInvestment(db.GetDB())

	loanServiceObj := service.NewLoan(
		loanStoreObj,
		userStoreObj,
		employeeStoreObj,
		investorStoreObj,
		investmentStoreObj,
		documentStoreObj,
	)

	v1 := router.Group("v1")
	{
		loanGroup := v1.Group("loan")
		{
			loanControllerObj := controller.NewLoan(loanServiceObj)
			loanGroup.POST("", loanControllerObj.Create)
			loanGroup.GET("/:id", loanControllerObj.Retrieve)

			// TODO: Add separate auth for different user group
			loanGroup.POST("/:id/approve", loanControllerObj.Approve)
			loanGroup.POST("/:id/invest", loanControllerObj.Invest)
		}
	}
	return router
}
