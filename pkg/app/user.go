package app

import "github.com/mateuszjurus/prometheus/pkg/domain"

type userSvc struct {
	DB domain.UserDB
}

func NewUserSvc(db domain.UserDB) domain.UserSvc {
	return userSvc{
		DB: db,
	}
}

func (us userSvc) GetUser(id int) (*domain.User, error) {
	return us.DB.GetUser(id)

}

func (us userSvc) ListUser(category string) ([]*domain.User, error) {
	return us.DB.ListUser(category)
}

func (us userSvc) CreateUser(user *domain.User) error {
	return us.DB.CreateUser(user)
}

func (us userSvc) DeleteUser(id int) error {
	return us.DB.DeleteUser(id)
}
