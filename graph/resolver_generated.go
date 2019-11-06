package graph

import (
	"context"
	"desafio-graph/src/resolvers"
	"desafio-graph/src/structs"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Sites(ctx context.Context, site *string) ([]*structs.Site, error) {
	return resolvers.GetSite(ctx, site)
}
