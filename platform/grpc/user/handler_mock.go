package user

import (
	"github.com/stretchr/testify/mock"
	Session "go-fiber-v2/pkg/libs/session"
)

type MockserverContext struct {
	mock.Mock
}

func (m *MockserverContext) TokenValidation(session *Session.Session, token string) (string, error) {
	args := m.Called(token)
	return args.Get(0).(string), args.Error(1)
}
