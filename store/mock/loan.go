// Code generated by MockGen. DO NOT EDIT.
// Source: loan.go

// Package store is a generated GoMock package.
package store

import (
	context "context"
	reflect "reflect"

	entity "github.com/1991-bishnu/loan-service/db/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockLoan is a mock of Loan interface.
type MockLoan struct {
	ctrl     *gomock.Controller
	recorder *MockLoanMockRecorder
}

// MockLoanMockRecorder is the mock recorder for MockLoan.
type MockLoanMockRecorder struct {
	mock *MockLoan
}

// NewMockLoan creates a new mock instance.
func NewMockLoan(ctrl *gomock.Controller) *MockLoan {
	mock := &MockLoan{ctrl: ctrl}
	mock.recorder = &MockLoanMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoan) EXPECT() *MockLoanMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockLoan) GetByID(ctx context.Context, id string) (*entity.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*entity.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockLoanMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockLoan)(nil).GetByID), ctx, id)
}

// GetByIDAndUserID mocks base method.
func (m *MockLoan) GetByIDAndUserID(ctx context.Context, id, userID string) (*entity.Loan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDAndUserID", ctx, id, userID)
	ret0, _ := ret[0].(*entity.Loan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDAndUserID indicates an expected call of GetByIDAndUserID.
func (mr *MockLoanMockRecorder) GetByIDAndUserID(ctx, id, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDAndUserID", reflect.TypeOf((*MockLoan)(nil).GetByIDAndUserID), ctx, id, userID)
}

// Insert mocks base method.
func (m *MockLoan) Insert(ctx context.Context, loan *entity.Loan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, loan)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockLoanMockRecorder) Insert(ctx, loan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockLoan)(nil).Insert), ctx, loan)
}

// Update mocks base method.
func (m *MockLoan) Update(ctx context.Context, loan *entity.Loan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, loan)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockLoanMockRecorder) Update(ctx, loan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLoan)(nil).Update), ctx, loan)
}
