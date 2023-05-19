package graph

import (
	"fmt"

	"github.com/2zqa/duwallpaper-graphql/graph/model"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	categories []*model.Category
}

func (r *Resolver) getCategory(id string) (*model.Category, error) {
	for _, c := range r.categories {
		if c.ID == id {
			return c, nil
		}
	}
	return &model.Category{}, fmt.Errorf("category with id %s does not exist", id)
}
