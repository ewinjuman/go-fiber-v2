package user

import (
	"github.com/stretchr/testify/mock"
)

type MockUserGrpc struct {
	mock.Mock
}

func (m *MockUserGrpc) TokenValidation(token string) (string, error) {
	args := m.Called(token)
	println(args.Get(0).(string))
	return args.Get(0).(string), args.Error(1)
}
