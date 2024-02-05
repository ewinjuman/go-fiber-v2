package queries

import (
	"github.com/stretchr/testify/mock"
	"go-fiber-v2/app/models"
)

type MockUserQueries struct {
	mock.Mock
}

func (m *MockUserQueries) InsertOneItem(req *models.User) (user *models.User, err error) {
	args := m.Called(req)
	return args.Get(0).(*models.User), args.Error(1)
}
