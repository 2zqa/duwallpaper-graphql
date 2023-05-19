package graph

import (
	"fmt"
	"strconv"

	"github.com/2zqa/duwallpaper-graphql/graph/model"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	wallpapers []*model.Wallpaper
	categories []*model.Category
}

func isNotInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err != nil
}

func (r *Resolver) getCategory(id string) (*model.Category, error) {
	if isNotInteger(id) {
		return &model.Category{}, fmt.Errorf("id must be an integer")
	}

	for _, c := range r.categories {
		if c.ID == id {
			return c, nil
		}
	}
	return &model.Category{}, fmt.Errorf("category with id %s does not exist", id)
}

func (r *Resolver) getWallpaper(id string) (*model.Wallpaper, error) {
	if isNotInteger(id) {
		return &model.Wallpaper{}, fmt.Errorf("id must be an integer")
	}

	for _, w := range r.wallpapers {
		if w.ID == id {
			return w, nil
		}
	}
	return &model.Wallpaper{}, fmt.Errorf("wallpaper with id %s does not exist", id)
}
