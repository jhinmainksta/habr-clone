// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/jhinmainksta/habr-clone/graph/my_model"
)

type Mutation struct {
}

type NewComment struct {
	Content  string `json:"content"`
	PostID   int    `json:"postID"`
	ParentID *int   `json:"parentID,omitempty"`
}

type NewPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Blocked *bool  `json:"blocked,omitempty"`
}

type Post struct {
	ID       string              `json:"id"`
	Title    string              `json:"title"`
	Content  string              `json:"content"`
	Blocked  *bool               `json:"blocked,omitempty"`
	Comments []*my_model.Comment `json:"comments"`
}

type Query struct {
}

type Subscription struct {
}
