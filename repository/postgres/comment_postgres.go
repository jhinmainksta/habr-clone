package postgres

import (
	"github.com/jhinmainksta/habr-clone/graph/model"
	"github.com/jhinmainksta/habr-clone/graph/my_model"
	"gorm.io/gorm"
)

type CommentPostgres struct {
	db *gorm.DB
}

func NewCommentPostgres(db *gorm.DB) *CommentPostgres {
	return &CommentPostgres{db: db}
}

func (h *CommentPostgres) CreateComment(NewComment model.NewComment) (*my_model.Comment, error) {
	if NewComment.ParentID == nil {
		id := 0
		NewComment.ParentID = &id
	}
	err := h.db.Table(commentsTable).Create(NewComment).Error

	comment := &my_model.Comment{
		Content:  NewComment.Content,
		PostID:   NewComment.PostID,
		UserID:   NewComment.UserID,
		ParentID: NewComment.ParentID,
	}

	return comment, err
}

func (h *CommentPostgres) Comment(id string) (*my_model.Comment, error) {
	Comment := &my_model.Comment{}
	err := h.db.First(Comment, id).Error

	return Comment, err
}

func (h *CommentPostgres) Comments(limit int, offset int) ([]*my_model.Comment, error) {
	var Comments []*my_model.Comment

	err := h.db.Limit(limit).Offset(offset).Find(&Comments).Error

	return Comments, err
}

func (h *CommentPostgres) CommentsComments(obj *my_model.Comment, limit int, offset int) ([]*my_model.Comment, error) {

	var Comments []*my_model.Comment

	err := h.db.Limit(limit).Offset(offset).Where("parent_id = ?", obj.ID).Find(&Comments).Error

	return Comments, err
}
