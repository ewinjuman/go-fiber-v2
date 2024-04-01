package queries

import (
	"github.com/stretchr/testify/mock"
	"go-fiber-v2/app/domain/entities"
)

type MockUserQueries struct {
	mock.Mock
}

func (m *MockUserQueries) InsertOneItem(req *entities.User) (user *entities.User, err error) {
	args := m.Called(req)
	return args.Get(0).(*entities.User), args.Error(1)
}
