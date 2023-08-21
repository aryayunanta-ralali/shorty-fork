// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repositories/contract.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	entity "github.com/aryayunanta-ralali/shorty/internal/entity"
	repositories "github.com/aryayunanta-ralali/shorty/internal/repositories"
	gomock "github.com/golang/mock/gomock"
)

// MockDBTransaction is a mock of DBTransaction interface.
type MockDBTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockDBTransactionMockRecorder
}

// MockDBTransactionMockRecorder is the mock recorder for MockDBTransaction.
type MockDBTransactionMockRecorder struct {
	mock *MockDBTransaction
}

// NewMockDBTransaction creates a new mock instance.
func NewMockDBTransaction(ctrl *gomock.Controller) *MockDBTransaction {
	mock := &MockDBTransaction{ctrl: ctrl}
	mock.recorder = &MockDBTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBTransaction) EXPECT() *MockDBTransactionMockRecorder {
	return m.recorder
}

// ExecTX mocks base method.
func (m *MockDBTransaction) ExecTX(ctx context.Context, options *sql.TxOptions, fn func(context.Context, repositories.StoreTX) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecTX", ctx, options, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecTX indicates an expected call of ExecTX.
func (mr *MockDBTransactionMockRecorder) ExecTX(ctx, options, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecTX", reflect.TypeOf((*MockDBTransaction)(nil).ExecTX), ctx, options, fn)
}

// MockStoreTX is a mock of StoreTX interface.
type MockStoreTX struct {
	ctrl     *gomock.Controller
	recorder *MockStoreTXMockRecorder
}

// MockStoreTXMockRecorder is the mock recorder for MockStoreTX.
type MockStoreTXMockRecorder struct {
	mock *MockStoreTX
}

// NewMockStoreTX creates a new mock instance.
func NewMockStoreTX(ctrl *gomock.Controller) *MockStoreTX {
	mock := &MockStoreTX{ctrl: ctrl}
	mock.recorder = &MockStoreTXMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreTX) EXPECT() *MockStoreTXMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockStoreTX) Execute(ctx context.Context, query string, values ...interface{}) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Execute", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockStoreTXMockRecorder) Execute(ctx, query interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockStoreTX)(nil).Execute), varargs...)
}

// Store mocks base method.
func (m *MockStoreTX) Store(ctx context.Context, tableName string, entity interface{}) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, tableName, entity)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store.
func (mr *MockStoreTXMockRecorder) Store(ctx, tableName, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockStoreTX)(nil).Store), ctx, tableName, entity)
}

// StoreBulk mocks base method.
func (m *MockStoreTX) StoreBulk(ctx context.Context, tableName string, entity interface{}) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreBulk", ctx, tableName, entity)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreBulk indicates an expected call of StoreBulk.
func (mr *MockStoreTXMockRecorder) StoreBulk(ctx, tableName, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreBulk", reflect.TypeOf((*MockStoreTX)(nil).StoreBulk), ctx, tableName, entity)
}

// Update mocks base method.
func (m *MockStoreTX) Update(ctx context.Context, tableName string, entity interface{}, whereConditions []repositories.WhereCondition) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, tableName, entity, whereConditions)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockStoreTXMockRecorder) Update(ctx, tableName, entity, whereConditions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStoreTX)(nil).Update), ctx, tableName, entity, whereConditions)
}

// Upsert mocks base method.
func (m *MockStoreTX) Upsert(ctx context.Context, tableName string, entity interface{}, onUpdate []repositories.WhereCondition) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, tableName, entity, onUpdate)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert.
func (mr *MockStoreTXMockRecorder) Upsert(ctx, tableName, entity, onUpdate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockStoreTX)(nil).Upsert), ctx, tableName, entity, onUpdate)
}

// MockShortUrls is a mock of ShortUrls interface.
type MockShortUrls struct {
	ctrl     *gomock.Controller
	recorder *MockShortUrlsMockRecorder
}

// MockShortUrlsMockRecorder is the mock recorder for MockShortUrls.
type MockShortUrlsMockRecorder struct {
	mock *MockShortUrls
}

// NewMockShortUrls creates a new mock instance.
func NewMockShortUrls(ctrl *gomock.Controller) *MockShortUrls {
	mock := &MockShortUrls{ctrl: ctrl}
	mock.recorder = &MockShortUrlsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShortUrls) EXPECT() *MockShortUrlsMockRecorder {
	return m.recorder
}

// FindBy mocks base method.
func (m *MockShortUrls) FindBy(ctx context.Context, cri repositories.FindShortUrlsCriteria) ([]entity.ShortUrls, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBy", ctx, cri)
	ret0, _ := ret[0].([]entity.ShortUrls)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBy indicates an expected call of FindBy.
func (mr *MockShortUrlsMockRecorder) FindBy(ctx, cri interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBy", reflect.TypeOf((*MockShortUrls)(nil).FindBy), ctx, cri)
}

// IncrementViewCount mocks base method.
func (m *MockShortUrls) IncrementViewCount(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementViewCount", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementViewCount indicates an expected call of IncrementViewCount.
func (mr *MockShortUrlsMockRecorder) IncrementViewCount(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementViewCount", reflect.TypeOf((*MockShortUrls)(nil).IncrementViewCount), ctx, id)
}

// Insert mocks base method.
func (m *MockShortUrls) Insert(ctx context.Context, data entity.ShortUrls) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockShortUrlsMockRecorder) Insert(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockShortUrls)(nil).Insert), ctx, data)
}

// Update mocks base method.
func (m *MockShortUrls) Update(ctx context.Context, data entity.ShortUrls) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockShortUrlsMockRecorder) Update(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockShortUrls)(nil).Update), ctx, data)
}
