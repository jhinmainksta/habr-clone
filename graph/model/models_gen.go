// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Postid  int    `json:"postid"`
	Userid  int    `json:"userid"`
	Parent  *int   `json:"parent,omitempty"`
}

type Mutation struct {
}

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Userid  int    `json:"userid"`
	Blocked bool   `json:"blocked"`
}

type Query struct {
}

type Subscription struct {
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
