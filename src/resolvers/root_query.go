package resolvers

import "desafio-graph/src/libs"

// QueryResolver represents root resolver
type QueryResolver struct {
	crawler  libs.Crawler
	monetary libs.Monetary
}

// NewQueryResolver instance a new QueryResolver
func NewQueryResolver(crawler libs.Crawler, monetary libs.Monetary) *QueryResolver {
	return &QueryResolver{crawler: crawler, monetary: monetary}
}
