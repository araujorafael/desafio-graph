package graph

import (
	"context"
	"desafio-graph/src/libs"
	"desafio-graph/src/resolvers"
	"desafio-graph/src/structs"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver is a implementation layer
type Resolver struct {
	sitesResolver *resolvers.SiteResolver
}

// NewResolver Create implementation and injects dependencies
func NewResolver(crawler libs.Crawler) *Resolver {
	return &Resolver{
		sitesResolver: resolvers.NewSiteResolver(crawler),
	}
}

// Query default query resolver implementation
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

// Sites is a resolver implementation to Site schema
func (r *queryResolver) Sites(ctx context.Context, site *string) ([]*structs.Site, error) {
	return r.sitesResolver.GetSite(ctx, site)
}
