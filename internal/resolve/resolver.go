package resolve

import "github.com/lodthe/graphql-example/internal/match"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	repo match.Repository
}

func NewResolver(r match.Repository) *Resolver {
	return &Resolver{repo: r}
}
