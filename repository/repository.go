package repository

import (
	"github.com/jhinmainksta/habr-clone/graph/model"
	"github.com/jhinmainksta/habr-clone/graph/my_model"
	"github.com/jhinmainksta/habr-clone/repository/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	HCPost
	HCComment
	HCUser
}

type HCPost interface {
	CreatePost(model.NewPost) (*model.Post, error)
	Post(string) (*model.Post, error)
	Posts(int, int) ([]*model.Post, error)
	PostsComments(*model.Post, int, int) ([]*my_model.Comment, error)
	BlockComments(string) (*model.Post, error)
}

type HCComment interface {
	CreateComment(model.NewComment) (*my_model.Comment, error)
	Comment(string) (*my_model.Comment, error)
	Comments(int, int) ([]*my_model.Comment, error)
	CommentsComments(*my_model.Comment, int, int) ([]*my_model.Comment, error)
}

type HCUser interface {
	CreateUser(model.NewUser) (*model.User, error)
	User(string) (*model.User, error)
	Users() ([]*model.User, error)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		HCPost:    postgres.NewPostPostgres(db),
		HCUser:    postgres.NewUserPostgres(db),
		HCComment: postgres.NewCommentPostgres(db),
	}
}
