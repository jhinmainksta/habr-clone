package repository

import (
	"github.com/jhinmainksta/habr-clone/graph/model"
	"gorm.io/gorm"
)

type Repository struct {
	HabrClone
}

type HabrClone interface {
	CreateUser(model.NewUser) (*model.User, error)
	CreatePost(model.NewPost) (*model.Post, error)
	CreateComment(model.NewComment) (*model.Comment, error)
	User(id string) (*model.User, error)
	Users() ([]*model.User, error)
	Post(id string) (*model.Post, error)
	Posts() ([]*model.Post, error)
	Comment(id string) (*model.Comment, error)
	Comments() ([]*model.Comment, error)
	PostsComments(obj *model.Post) ([]*model.Comment, error)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		HabrClone: NewHabrClonePG(db),
	}
}
