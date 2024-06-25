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
	User(string) (*model.User, error)
	Users() ([]*model.User, error)
	Post(string) (*model.Post, error)
	Posts(int, int) ([]*model.Post, error)
	Comment(string) (*model.Comment, error)
	Comments(int, int) ([]*model.Comment, error)
	PostsComments(*model.Post, int, int) ([]*model.Comment, error)
	CommentsComments(*model.Comment, int, int) ([]*model.Comment, error)
	BlockComments(string) (*model.Post, error)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		HabrClone: NewHabrClonePG(db),
	}
}
