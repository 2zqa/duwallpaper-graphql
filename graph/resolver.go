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
	tags       []*model.Tag
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

func (r *Resolver) getTag(id string) (*model.Tag, error) {
	if isNotInteger(id) {
		return &model.Tag{}, fmt.Errorf("id must be an integer")
	}

	for _, t := range r.tags {
		if t.ID == id {
			return t, nil
		}
	}
	return &model.Tag{}, fmt.Errorf("tag with id %s does not exist", id)
}

func (r *Resolver) deleteTag(id string) (bool, error) {
	if isNotInteger(id) {
		return false, fmt.Errorf("id must be an integer")
	}

	for i, t := range r.tags {
		if t.ID == id {
			// Efficiently delete without preserving order
			r.tags[i] = r.tags[len(r.tags)-1]
			return true, nil
		}
	}
	return false, fmt.Errorf("tag with id %s does not exist", id)
}
