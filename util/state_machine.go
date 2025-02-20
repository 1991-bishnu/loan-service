package util

import (
	"github.com/1991-bishnu/loan-service/constant"
	"github.com/1991-bishnu/loan-service/custom_error"
)

func GetNextStage(currentStatus string) (string, error) {
	switch currentStatus {
	case constant.LoanStatusProposed:
		return constant.LoanStatusApproved, nil
	case constant.LoanStatusApproved:
		return constant.LoanStatusInvested, nil
	case constant.LoanStatusInvested:
		return constant.LoanStatusDisbursed, nil
	default:
		return "", custom_error.ErrInvalidTransition
	}
}
