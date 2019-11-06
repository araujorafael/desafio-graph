package resolvers

import (
	"context"
	"desafio-graph/src/structs"

	"github.com/vektah/gqlparser/gqlerror"
)

// GetSite return site info
func GetSite(ctx context.Context, site *string) ([]*structs.Site, error) {
	if *site != "https://www.smartmei.com.br" {
		return []*structs.Site{}, gqlerror.Errorf("404 - NOT FOUND")
	}

	found := []*structs.Site{{URL: "https://www.smartmei.com.br"}}
	return found, nil
}
