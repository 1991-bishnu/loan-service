package controller

import (
	"errors"
	"net/http"

	"github.com/1991-bishnu/loan-service/custom_error"
	"github.com/1991-bishnu/loan-service/model"
	"github.com/1991-bishnu/loan-service/service"
	"github.com/gin-gonic/gin"
)

type Loan interface {
	Create(c *gin.Context)
	Retrieve(c *gin.Context)
	Approve(c *gin.Context)
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
	if errors.Is(err, custom_error.UserNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
	if errors.Is(err, custom_error.LoanNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)

}

func (obj *loan) Approve(c *gin.Context) {
	ctx := c.Request.Context()

	req := &model.ApproveLoanReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.LoanID = c.Param("id")
	if req.LoanID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "loan id is missing in URL"})
		return
	}

	err := obj.service.Approve(ctx, req)
	if err != nil {
		errorMap := map[error]int{
			custom_error.LoanNotFound:         http.StatusBadRequest,
			custom_error.EmployeeNotFound:     http.StatusBadRequest,
			custom_error.ErrInvalidTransition: http.StatusBadRequest,
		}
		for key, status := range errorMap {
			if errors.Is(err, key) {
				c.JSON(status, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusAccepted)
}
