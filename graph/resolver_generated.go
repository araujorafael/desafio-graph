package graph

import (
	"desafio-graph/src/libs"
	"desafio-graph/src/resolvers"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	crawler  libs.Crawler
	monetary libs.Monetary
}

func NewResolver(crawler libs.Crawler, monetary libs.MonetaryImpl) Config {
	return Config{Resolvers: &Resolver{crawler, monetary}}
}

func (r *Resolver) Query() QueryResolver {
	return resolvers.NewQueryResolver(r.crawler, r.monetary)
}
