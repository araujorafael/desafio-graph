package resolvers

import (
	"context"
	"desafio-graph/src/structs"
	"fmt"
	"strconv"

	"github.com/vektah/gqlparser/gqlerror"
)

var currencySigns = map[string]string{
	"BRL": "R$",
	"USD": "$",
	"EUR": "€",
}

// Sites return site info
func (qr *QueryResolver) Sites(ctx context.Context, site *string) ([]*structs.Site, error) {
	if *site != "https://www.smartmei.com.br" {
		return []*structs.Site{}, gqlerror.Errorf("404 - RESOURCE NOT FOUND")
	}

	pricing := qr.crawler.GetPricing(*site)

	charges := make([]*structs.Charge, len(pricing))
	for i := range pricing {
		charge := &structs.Charge{
			Name:     pricing[i].Title,
			BrlValue: fmt.Sprintf("%s: %s", currencySigns["BRL"], pricing[i].NormalValue),
			UsdValue: qr.convertValue(pricing[i].NormalValue, "USD"),
			EurValue: qr.convertValue(pricing[i].NormalValue, "EUR"),
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

func (qr *QueryResolver) convertValue(value string, currency string) string {
	if value == "--" {
		return "Não disponvel para este plano"
	}

	i, _ := strconv.ParseFloat(value, 64)

	priceTag := qr.monetary.Convert(i, currency)
	return fmt.Sprintf("%s: %.2f", currencySigns[currency], priceTag)
}
