package queries

import (
	"gitlab.pede.id/otto-library/golang/share-pkg/session"
	"go-fiber-v2/app/domain/entities"
	"go-fiber-v2/pkg/repository"
	"go-fiber-v2/platform/database"
)

type (
	UserQueriesService interface {
		InsertOneItem(req *entities.User) (*entities.User, error)
	}

	userQueries struct {
		session *session.Session
	}
)

func NewUserQueries(session *session.Session) (rep UserQueriesService) {
	return &userQueries{session: session}
}

func (r *userQueries) InsertOneItem(req *entities.User) (user *entities.User, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	newUser := new(entities.User)
	err = db.Omit("updated_at").Create(req).Scan(newUser).Error
	if err != nil {
		err = repository.ConvMysqlErr(err)
		return
	}
	return newUser, nil
}
