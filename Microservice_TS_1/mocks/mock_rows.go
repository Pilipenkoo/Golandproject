package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockRows struct {
	mock.Mock
}

func (m *MockRows) Next() bool {
	ret := m.Called()
	return ret.Bool(0)
}

func (m *MockRows) Scan(dest ...interface{}) error {
	args := m.Called(dest...)
	return args.Error(0)
}

func (m *MockRows) Err() error {
	ret := m.Called()
	return ret.Error(0)
}

func (m *MockRows) Close() {}
