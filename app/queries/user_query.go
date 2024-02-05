package queries

import (
	"go-fiber-v2/app/models"
	"go-fiber-v2/pkg/libs/session"
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
	row := new(models.User)
	err = db.Omit("updated_at").Create(req).Scan(row).Error

	if err != nil {
		return
	}
	user = row

	return
}
