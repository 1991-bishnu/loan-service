package controller

import (
	"net/http"

	"github.com/1991-bishnu/loan-service/model"
	"github.com/1991-bishnu/loan-service/service"
	"github.com/gin-gonic/gin"
)

type Loan interface {
	Create(c *gin.Context)
	Retrieve(c *gin.Context)
}
type loan struct {
	service service.Loan
}

func NewLoan(service service.Loan) Loan {
	return &loan{
		service: service,
	}

}

func (obj *loan) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req model.CreateLoanReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := obj.service.Create(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)

}

func (obj *loan) Retrieve(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "loan id is missing in URL"})
		return
	}

	req := &model.RetrieveLoanReq{
		LoanID: id,
		UserID: c.Query("user_id"),
	}

	res, err := obj.service.Retrieve(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)

}
