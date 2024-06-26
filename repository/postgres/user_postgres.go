package postgres

import (
	"github.com/jhinmainksta/habr-clone/graph/model"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (h *UserPostgres) CreateUser(NewUser model.NewUser) (*model.User, error) {
	err := h.db.Table(usersTable).Create(NewUser).Error

	user := &model.User{
		Username: NewUser.Username,
		Password: NewUser.Password,
	}

	return user, err
}

func (h *UserPostgres) User(id string) (*model.User, error) {
	user := &model.User{}
	err := h.db.First(user, id).Error

	return user, err
}

func (h *UserPostgres) Users() ([]*model.User, error) {
	var users []*model.User

	err := h.db.Find(&users).Error

	return users, err
}
