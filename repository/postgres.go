package repository

import (
	"fmt"

	"github.com/jhinmainksta/habr-clone/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	postsTable    = "posts"
	commentsTable = "comments"
	usersTable    = "users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}

type HabrClonePG struct {
	db *gorm.DB
}

func NewHabrClonePG(db *gorm.DB) *HabrClonePG {
	return &HabrClonePG{db: db}
}

func (h *HabrClonePG) CreateUser(NewUser model.NewUser) (*model.User, error) {
	err := h.db.Table(usersTable).Create(NewUser).Error

	user := &model.User{
		Username: NewUser.Username,
		Password: NewUser.Password,
	}

	return user, err
}

func (h *HabrClonePG) User(id string) (*model.User, error) {
	user := &model.User{}
	err := h.db.First(user, id).Error

	return user, err
}

func (h *HabrClonePG) Users() ([]*model.User, error) {
	var users []*model.User

	err := h.db.Find(&users).Error

	return users, err
}

func (h *HabrClonePG) CreatePost(NewPost model.NewPost) (*model.Post, error) {

	post := &model.Post{
		Title:   NewPost.Title,
		Content: NewPost.Content,
		UserID:  NewPost.UserID,
		Blocked: NewPost.Blocked,
	}

	if NewPost.Blocked == nil {
		NewPost.Blocked = new(bool)
		post.Blocked = new(bool)
	}

	err := h.db.Table(postsTable).Create(NewPost).Error

	return post, err
}

func (h *HabrClonePG) Post(id string) (*model.Post, error) {
	post := &model.Post{}
	err := h.db.First(post, id).Error

	return post, err
}

func (h *HabrClonePG) Posts(limit int, offset int) ([]*model.Post, error) {
	var posts []*model.Post

	err := h.db.Limit(limit).Offset(offset).Find(&posts).Error

	return posts, err
}

func (h *HabrClonePG) CreateComment(NewComment model.NewComment) (*model.Comment, error) {
	err := h.db.Table(commentsTable).Create(NewComment).Error

	comment := &model.Comment{
		Content:  NewComment.Content,
		PostID:   NewComment.PostID,
		UserID:   NewComment.UserID,
		ParentID: NewComment.ParentID,
	}

	return comment, err
}

func (h *HabrClonePG) Comment(id string) (*model.Comment, error) {
	Comment := &model.Comment{}
	err := h.db.First(Comment, id).Error

	return Comment, err
}

func (h *HabrClonePG) Comments(limit int, offset int) ([]*model.Comment, error) {
	var Comments []*model.Comment

	err := h.db.Limit(limit).Offset(offset).Find(&Comments).Error

	return Comments, err
}

func (h *HabrClonePG) PostsComments(obj *model.Post, limit int, offset int) ([]*model.Comment, error) {
	var Comments []*model.Comment

	err := h.db.Limit(limit).Offset(offset).Where("post_id = ?", obj.ID).Find(&Comments).Error

	return Comments, err
}

func (h *HabrClonePG) CommentsComments(obj *model.Comment, limit int, offset int) ([]*model.Comment, error) {

	var Comments []*model.Comment

	err := h.db.Limit(limit).Offset(offset).Where("parent_id = ?", obj.ID).Find(&Comments).Error

	return Comments, err
}

func (h *HabrClonePG) BlockComments(postID string) (*model.Post, error) {

	post := &model.Post{ID: postID}
	err := h.db.Model(post).Update("blocked", true).Error

	if err == nil {
		err = h.db.First(post).Error
	}

	return post, err
}
