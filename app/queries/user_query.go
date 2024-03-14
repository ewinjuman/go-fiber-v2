package queries

import (
	"gitlab.pede.id/otto-library/golang/share-pkg/session"
	"go-fiber-v2/app/models"
	"go-fiber-v2/platform/database"
)

type (
	UserQueriesService interface {
		InsertOneItem(req *models.User) (*models.User, error)
	}

	userQueries struct {
		session *session.Session
	}
)

func NewUserQueries(session *session.Session) (rep UserQueriesService) {
	return &userQueries{session: session}
}

func (r *userQueries) InsertOneItem(req *models.User) (user *models.User, err error) {
	db, err := database.MysqlConnection(r.session)
	if err != nil {
		return
	}
	newUser := new(models.User)
	err = db.Omit("updated_at").Create(req).Scan(newUser).Error
	if err != nil {
		return
	}
	return newUser, nil
}
