package resolvers

import (
	"context"
	"desafio-graph/src/libs"
	"desafio-graph/src/structs"

	"github.com/vektah/gqlparser/gqlerror"
)

// GetSite return site info
func GetSite(ctx context.Context, site *string) ([]*structs.Site, error) {
	if *site != "https://www.smartmei.com.br" {
		return []*structs.Site{}, gqlerror.Errorf("404 - NOT FOUND")
	}

	pricing := libs.GetPricing(*site)

	charges := make([]*structs.Charge, len(pricing))
	for i := range pricing {
		cg := &structs.Charge{
			Name:     pricing[i].Title,
			BrlValue: pricing[i].NormalValue,
		}

		charges[i] = cg
	}

	siteData := &structs.Site{
		URL:     *site,
		Charges: charges,
	}

	found := []*structs.Site{siteData}
	return found, nil
}
