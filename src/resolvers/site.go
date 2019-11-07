package resolvers

import (
	"context"
	"desafio-graph/src/libs"
	"desafio-graph/src/structs"

	"github.com/vektah/gqlparser/gqlerror"
)

// SiteResolver implementation
type SiteResolver struct {
	crawler libs.Crawler
}

// NewSiteResolver create SiteResolver implementation
func NewSiteResolver(crawler libs.Crawler) *SiteResolver {
	return &SiteResolver{
		crawler: crawler,
	}
}

// GetSite return site info
func (sr *SiteResolver) GetSite(ctx context.Context, site *string) ([]*structs.Site, error) {
	if *site != "https://www.smartmei.com.br" {
		return []*structs.Site{}, gqlerror.Errorf("404 - NOT FOUND")
	}

	pricing := sr.crawler.GetPricing(*site)

	charges := make([]*structs.Charge, len(pricing))
	for i := range pricing {
		charge := &structs.Charge{
			Name:     pricing[i].Title,
			BrlValue: pricing[i].NormalValue,
		}

		charges[i] = charge
	}

	siteData := &structs.Site{
		URL:     *site,
		Charges: charges,
	}

	found := []*structs.Site{siteData}
	return found, nil
}
