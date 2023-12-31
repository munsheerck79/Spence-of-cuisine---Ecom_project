// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interfacess/couponInterfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/munsheerck79/Ecom_project.git/pkg/domain"
	request "github.com/munsheerck79/Ecom_project.git/util/request"
)

// MockCouponService is a mock of CouponService interface.
type MockCouponService struct {
	ctrl     *gomock.Controller
	recorder *MockCouponServiceMockRecorder
}

// MockCouponServiceMockRecorder is the mock recorder for MockCouponService.
type MockCouponServiceMockRecorder struct {
	mock *MockCouponService
}

// NewMockCouponService creates a new mock instance.
func NewMockCouponService(ctrl *gomock.Controller) *MockCouponService {
	mock := &MockCouponService{ctrl: ctrl}
	mock.recorder = &MockCouponServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCouponService) EXPECT() *MockCouponServiceMockRecorder {
	return m.recorder
}

// AddCoupon mocks base method.
func (m *MockCouponService) AddCoupon(c context.Context, body domain.Coupon) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCoupon", c, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCoupon indicates an expected call of AddCoupon.
func (mr *MockCouponServiceMockRecorder) AddCoupon(c, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCoupon", reflect.TypeOf((*MockCouponService)(nil).AddCoupon), c, body)
}

// EditCoupon mocks base method.
func (m *MockCouponService) EditCoupon(c context.Context, body request.EditCoupon) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditCoupon", c, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditCoupon indicates an expected call of EditCoupon.
func (mr *MockCouponServiceMockRecorder) EditCoupon(c, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditCoupon", reflect.TypeOf((*MockCouponService)(nil).EditCoupon), c, body)
}

// GetCoupon mocks base method.
func (m *MockCouponService) GetCoupon(c context.Context) ([]domain.Coupon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCoupon", c)
	ret0, _ := ret[0].([]domain.Coupon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCoupon indicates an expected call of GetCoupon.
func (mr *MockCouponServiceMockRecorder) GetCoupon(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCoupon", reflect.TypeOf((*MockCouponService)(nil).GetCoupon), c)
}
