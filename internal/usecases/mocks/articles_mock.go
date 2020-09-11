// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nachogoca/golang-example-rest-api-layout/internal/usecases (interfaces: ArticlesStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entities "github.com/nachogoca/golang-example-rest-api-layout/internal/entities"
	reflect "reflect"
)

// MockArticlesStore is a mock of ArticlesStore interface
type MockArticlesStore struct {
	ctrl     *gomock.Controller
	recorder *MockArticlesStoreMockRecorder
}

// MockArticlesStoreMockRecorder is the mock recorder for MockArticlesStore
type MockArticlesStoreMockRecorder struct {
	mock *MockArticlesStore
}

// NewMockArticlesStore creates a new mock instance
func NewMockArticlesStore(ctrl *gomock.Controller) *MockArticlesStore {
	mock := &MockArticlesStore{ctrl: ctrl}
	mock.recorder = &MockArticlesStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockArticlesStore) EXPECT() *MockArticlesStoreMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockArticlesStore) Create(arg0 context.Context, arg1 entities.Article) (entities.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(entities.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockArticlesStoreMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArticlesStore)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockArticlesStore) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockArticlesStoreMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArticlesStore)(nil).Delete), arg0, arg1)
}

// GetAll mocks base method
func (m *MockArticlesStore) GetAll(arg0 context.Context) ([]entities.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]entities.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockArticlesStoreMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockArticlesStore)(nil).GetAll), arg0)
}

// GetOne mocks base method
func (m *MockArticlesStore) GetOne(arg0 context.Context, arg1 string) (entities.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", arg0, arg1)
	ret0, _ := ret[0].(entities.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne
func (mr *MockArticlesStoreMockRecorder) GetOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockArticlesStore)(nil).GetOne), arg0, arg1)
}

// Update mocks base method
func (m *MockArticlesStore) Update(arg0 context.Context, arg1 entities.Article) (entities.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(entities.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockArticlesStoreMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArticlesStore)(nil).Update), arg0, arg1)
}