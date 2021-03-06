// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nachogoca/golang-example-rest-api-layout/internal/transports (interfaces: ArticlesUsecase)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entities "github.com/nachogoca/golang-example-rest-api-layout/internal/entities"
	reflect "reflect"
)

// MockArticlesUsecase is a mock of ArticlesUsecase interface
type MockArticlesUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockArticlesUsecaseMockRecorder
}

// MockArticlesUsecaseMockRecorder is the mock recorder for MockArticlesUsecase
type MockArticlesUsecaseMockRecorder struct {
	mock *MockArticlesUsecase
}

// NewMockArticlesUsecase creates a new mock instance
func NewMockArticlesUsecase(ctrl *gomock.Controller) *MockArticlesUsecase {
	mock := &MockArticlesUsecase{ctrl: ctrl}
	mock.recorder = &MockArticlesUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockArticlesUsecase) EXPECT() *MockArticlesUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockArticlesUsecase) Create(arg0 context.Context, arg1 entities.Article) (entities.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(entities.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockArticlesUsecaseMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArticlesUsecase)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockArticlesUsecase) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockArticlesUsecaseMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArticlesUsecase)(nil).Delete), arg0, arg1)
}

// GetAll mocks base method
func (m *MockArticlesUsecase) GetAll(arg0 context.Context) ([]entities.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]entities.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockArticlesUsecaseMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockArticlesUsecase)(nil).GetAll), arg0)
}

// GetOne mocks base method
func (m *MockArticlesUsecase) GetOne(arg0 context.Context, arg1 string) (entities.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", arg0, arg1)
	ret0, _ := ret[0].(entities.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne
func (mr *MockArticlesUsecaseMockRecorder) GetOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockArticlesUsecase)(nil).GetOne), arg0, arg1)
}

// Update mocks base method
func (m *MockArticlesUsecase) Update(arg0 context.Context, arg1 entities.Article) (entities.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(entities.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockArticlesUsecaseMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArticlesUsecase)(nil).Update), arg0, arg1)
}
