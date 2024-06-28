package graph

import (
	"sync"

	"github.com/jhinmainksta/habr-clone/graph/my_model"
	"github.com/jhinmainksta/habr-clone/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	repo   *repository.Repository
	subs   map[string]map[string]chan *my_model.Comment
	mu     sync.Mutex
	limit  int
	offset int
}

func NewResolver(repo *repository.Repository, subs map[string]map[string]chan *my_model.Comment, limit int, offset int) *Resolver {
	return &Resolver{repo: repo, subs: subs, limit: limit, offset: offset}
}
