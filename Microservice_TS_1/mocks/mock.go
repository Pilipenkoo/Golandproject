package mocks

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockDB) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	//TODO implement me
	panic("implement me")
}

func (m *MockDB) LargeObjects() pgx.LargeObjects {
	//TODO implement me
	panic("implement me")
}

func (m *MockDB) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockDB) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockDB) Conn() *pgx.Conn {
	//TODO implement me
	panic("implement me")
}

func (m *MockDB) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	args = append([]interface{}{ctx, sql}, args...)
	ret := m.Called(args...)
	rows := ret.Get(0).(pgx.Rows)
	err := ret.Error(1)
	return rows, err
}

func (m *MockDB) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	args = append([]interface{}{ctx, sql}, args...)
	ret := m.Called(args...)
	commandTag := ret.Get(0).(pgconn.CommandTag)
	err := ret.Error(1)
	return commandTag, err
}

func (m *MockDB) Begin(ctx context.Context) (pgx.Tx, error) {
	ret := m.Called(ctx)
	tx := ret.Get(0).(pgx.Tx)
	err := ret.Error(1)
	return tx, err
}

func (m *MockDB) BeginFunc(ctx context.Context, f func(pgx.Tx) error) error {
	ret := m.Called(ctx, f)
	return ret.Error(0)
}

func (m *MockDB) Commit(ctx context.Context) error {
	ret := m.Called(ctx)
	return ret.Error(0)
}

func (m *MockDB) Rollback(ctx context.Context) error {
	ret := m.Called(ctx)
	return ret.Error(0)
}

func (m *MockDB) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	args = append([]interface{}{ctx, sql}, args...)
	ret := m.Called(args...)
	return ret.Get(0).(pgx.Row)
}
