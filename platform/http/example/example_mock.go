package example

import (
	"github.com/stretchr/testify/mock"
)

type MockUserHttp struct {
	mock.Mock
}

func (m *MockUserHttp) TokenSessionValidation(request ValidateSessionRequest) (response ValidateSessionResponse, err error) {
	args := m.Called(request)
	return args.Get(0).(ValidateSessionResponse), args.Error(1)
}
