package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/jhinmainksta/habr-clone/graph/model"
)

// Comments is the resolver for the comments field.
func (r *commentResolver) Comments(ctx context.Context, obj *model.Comment, limit *int, offset *int) ([]*model.Comment, error) {
	if limit == nil {
		limit = &r.limit
	}

	if offset == nil {
		offset = &r.offset
	}

	return r.repo.CommentsComments(obj, *limit, *offset)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.repo.CreateUser(input)
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	return r.repo.CreatePost(input)
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	return r.repo.CreateComment(input)
}

// BlockComments is the resolver for the blockComments field.
func (r *mutationResolver) BlockComments(ctx context.Context, postID string) (*model.Post, error) {
	return r.repo.BlockComments(postID)
}

// Comments is the resolver for the comments field.
func (r *postResolver) Comments(ctx context.Context, obj *model.Post, limit *int, offset *int) ([]*model.Comment, error) {
	if limit == nil {
		limit = &r.limit
	}

	if offset == nil {
		offset = &r.offset
	}

	return r.repo.PostsComments(obj, *limit, *offset)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.repo.User(id)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.repo.Users()
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, limit *int, offset *int) ([]*model.Post, error) {
	if limit == nil {
		limit = &r.limit
	}

	if offset == nil {
		offset = &r.offset
	}

	return r.repo.Posts(*limit, *offset)
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	return r.repo.Post(id)
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, limit *int, offset *int) ([]*model.Comment, error) {
	if limit == nil {
		limit = &r.limit
	}

	if offset == nil {
		offset = &r.offset
	}
	return r.repo.Comments(*limit, *offset)
}

// Comment is the resolver for the comment field.
func (r *queryResolver) Comment(ctx context.Context, id string) (*model.Comment, error) {
	return r.repo.Comment(id)
}

// CommentAdded is the resolver for the commentAdded field.
func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	panic(fmt.Errorf("not implemented: CommentAdded - commentAdded"))
}

// Comment returns CommentResolver implementation.
func (r *Resolver) Comment() CommentResolver { return &commentResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Post returns PostResolver implementation.
func (r *Resolver) Post() PostResolver { return &postResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type commentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
