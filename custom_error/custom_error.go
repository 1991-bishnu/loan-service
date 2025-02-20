package custom_error

import (
	"errors"
)

var LoanNotFound = errors.New("loan not found")
var UserNotFound = errors.New("user not found")
var EmployeeNotFound = errors.New("employee not found")
var ErrInvalidTransition = errors.New("invalid status transition")
