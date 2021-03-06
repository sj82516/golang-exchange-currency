// Code generated by MockGen. DO NOT EDIT.
// Source: exchange_price_model.go

// Package src is a generated GoMock package.
package src

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIExchangePriceModel is a mock of IExchangePriceModel interface.
type MockIExchangePriceModel struct {
	ctrl     *gomock.Controller
	recorder *MockIExchangePriceModelMockRecorder
}

// MockIExchangePriceModelMockRecorder is the mock recorder for MockIExchangePriceModel.
type MockIExchangePriceModelMockRecorder struct {
	mock *MockIExchangePriceModel
}

// NewMockIExchangePriceModel creates a new mock instance.
func NewMockIExchangePriceModel(ctrl *gomock.Controller) *MockIExchangePriceModel {
	mock := &MockIExchangePriceModel{ctrl: ctrl}
	mock.recorder = &MockIExchangePriceModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIExchangePriceModel) EXPECT() *MockIExchangePriceModelMockRecorder {
	return m.recorder
}

// GetExchangeRate mocks base method.
func (m *MockIExchangePriceModel) GetExchangeRate(arg0, arg1 string, arg2 chan<- ExchangeRateResult) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetExchangeRate", arg0, arg1, arg2)
}

// GetExchangeRate indicates an expected call of GetExchangeRate.
func (mr *MockIExchangePriceModelMockRecorder) GetExchangeRate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRate", reflect.TypeOf((*MockIExchangePriceModel)(nil).GetExchangeRate), arg0, arg1, arg2)
}
