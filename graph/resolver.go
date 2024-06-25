package graph

import "github.com/jhinmainksta/habr-clone/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	repo   *repository.Repository
	limit  int
	offset int
}

func NewResolver(repo *repository.Repository, limit int, offset int) *Resolver {
	return &Resolver{repo: repo, limit: limit, offset: offset}
}
