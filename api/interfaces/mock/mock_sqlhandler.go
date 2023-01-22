// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces/database/sqlhandler.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	database "github.com/kory-jp/vue_go/api/interfaces/database"
)

// MockSqlHandler is a mock of SqlHandler interface.
type MockSqlHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSqlHandlerMockRecorder
}

// MockSqlHandlerMockRecorder is the mock recorder for MockSqlHandler.
type MockSqlHandlerMockRecorder struct {
	mock *MockSqlHandler
}

// NewMockSqlHandler creates a new mock instance.
func NewMockSqlHandler(ctrl *gomock.Controller) *MockSqlHandler {
	mock := &MockSqlHandler{ctrl: ctrl}
	mock.recorder = &MockSqlHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSqlHandler) EXPECT() *MockSqlHandlerMockRecorder {
	return m.recorder
}

// DoInTx mocks base method.
func (m *MockSqlHandler) DoInTx(arg0 func(*sql.Tx) (interface{}, error)) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoInTx", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoInTx indicates an expected call of DoInTx.
func (mr *MockSqlHandlerMockRecorder) DoInTx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoInTx", reflect.TypeOf((*MockSqlHandler)(nil).DoInTx), arg0)
}

// Execute mocks base method.
func (m *MockSqlHandler) Execute(arg0 string, arg1 ...interface{}) (database.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Execute", varargs...)
	ret0, _ := ret[0].(database.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockSqlHandlerMockRecorder) Execute(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSqlHandler)(nil).Execute), varargs...)
}

// Query mocks base method.
func (m *MockSqlHandler) Query(arg0 string, arg1 ...interface{}) (database.Row, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(database.Row)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockSqlHandlerMockRecorder) Query(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockSqlHandler)(nil).Query), varargs...)
}

// TransExecute mocks base method.
func (m *MockSqlHandler) TransExecute(arg0 *sql.Tx, arg1 string, arg2 ...interface{}) (database.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TransExecute", varargs...)
	ret0, _ := ret[0].(database.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransExecute indicates an expected call of TransExecute.
func (mr *MockSqlHandlerMockRecorder) TransExecute(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransExecute", reflect.TypeOf((*MockSqlHandler)(nil).TransExecute), varargs...)
}

// MockResult is a mock of Result interface.
type MockResult struct {
	ctrl     *gomock.Controller
	recorder *MockResultMockRecorder
}

// MockResultMockRecorder is the mock recorder for MockResult.
type MockResultMockRecorder struct {
	mock *MockResult
}

// NewMockResult creates a new mock instance.
func NewMockResult(ctrl *gomock.Controller) *MockResult {
	mock := &MockResult{ctrl: ctrl}
	mock.recorder = &MockResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResult) EXPECT() *MockResultMockRecorder {
	return m.recorder
}

// LastInsertId mocks base method.
func (m *MockResult) LastInsertId() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastInsertId")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastInsertId indicates an expected call of LastInsertId.
func (mr *MockResultMockRecorder) LastInsertId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastInsertId", reflect.TypeOf((*MockResult)(nil).LastInsertId))
}

// RowsAffected mocks base method.
func (m *MockResult) RowsAffected() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RowsAffected")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RowsAffected indicates an expected call of RowsAffected.
func (mr *MockResultMockRecorder) RowsAffected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RowsAffected", reflect.TypeOf((*MockResult)(nil).RowsAffected))
}

// MockRow is a mock of Row interface.
type MockRow struct {
	ctrl     *gomock.Controller
	recorder *MockRowMockRecorder
}

// MockRowMockRecorder is the mock recorder for MockRow.
type MockRowMockRecorder struct {
	mock *MockRow
}

// NewMockRow creates a new mock instance.
func NewMockRow(ctrl *gomock.Controller) *MockRow {
	mock := &MockRow{ctrl: ctrl}
	mock.recorder = &MockRowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRow) EXPECT() *MockRowMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRow) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRowMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRow)(nil).Close))
}

// Err mocks base method.
func (m *MockRow) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockRowMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockRow)(nil).Err))
}

// Next mocks base method.
func (m *MockRow) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockRowMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockRow)(nil).Next))
}

// Scan mocks base method.
func (m *MockRow) Scan(arg0 ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockRowMockRecorder) Scan(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockRow)(nil).Scan), arg0...)
}