package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/2zqa/duwallpaper-graphql/graph/model"
)

// CreateWallpaper is the resolver for the createWallpaper field.
func (r *mutationResolver) CreateWallpaper(ctx context.Context, input model.NewWallpaper) (*model.Wallpaper, error) {
	panic(fmt.Errorf("not implemented: CreateWallpaper - createWallpaper"))
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, name string) (*model.Category, error) {
	id := len(r.categories) + 1
	newCategory := &model.Category{
		ID:         fmt.Sprintf("%d", id),
		Name:       name,
		Wallpapers: make([]*model.Wallpaper, 0),
	}
	r.categories = append(r.categories, newCategory)
	return newCategory, nil
}

// CreateTag is the resolver for the createTag field.
func (r *mutationResolver) CreateTag(ctx context.Context, name string) (bool, error) {
	panic(fmt.Errorf("not implemented: CreateTag - createTag"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// Wallpapers is the resolver for the wallpapers field.
func (r *queryResolver) Wallpapers(ctx context.Context) ([]*model.Wallpaper, error) {
	panic(fmt.Errorf("not implemented: Wallpapers - wallpapers"))
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.categories, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Wallpaper is the resolver for the wallpaper field.
func (r *queryResolver) Wallpaper(ctx context.Context, id string) (*model.Wallpaper, error) {
	panic(fmt.Errorf("not implemented: Wallpaper - wallpaper"))
}

// Category is the resolver for the category field.
func (r *queryResolver) Category(ctx context.Context, id string) (*model.Category, error) {
	return r.getCategory(id)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
