package postgres

import (
	"github.com/jhinmainksta/habr-clone/graph/model"
	"github.com/jhinmainksta/habr-clone/graph/my_model"
	"gorm.io/gorm"
)

type PostPostgres struct {
	db *gorm.DB
}

func NewPostPostgres(db *gorm.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (h *PostPostgres) CreatePost(NewPost model.NewPost) (*model.Post, error) {

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

func (h *PostPostgres) Post(id string) (*model.Post, error) {
	post := &model.Post{}
	err := h.db.First(post, id).Error

	return post, err
}

func (h *PostPostgres) Posts(limit int, offset int) ([]*model.Post, error) {
	var posts []*model.Post

	err := h.db.Limit(limit).Offset(offset).Find(&posts).Error

	return posts, err
}

func (h *PostPostgres) PostsComments(obj *model.Post, limit int, offset int) ([]*my_model.Comment, error) {

	Comments := []*my_model.Comment{}

	err := h.db.Limit(limit).Offset(offset).Where("post_id = ? AND parent_id = ?", obj.ID, 0).Find(&Comments).Error

	return Comments, err
}

func (h *PostPostgres) BlockComments(postID string) (*model.Post, error) {

	post := &model.Post{ID: postID}
	err := h.db.Model(post).Update("blocked", true).Error

	if err == nil {
		err = h.db.First(post).Error
	}

	return post, err
}
