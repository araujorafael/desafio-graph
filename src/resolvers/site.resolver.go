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

	basicPlanCharges := make([]*structs.Charge, len(pricing))
	professionalPlanCharges := make([]*structs.Charge, len(pricing))
	for i := range pricing {
		basic := &structs.Charge{
			Name:     pricing[i].Title,
			BrlValue: qr.convertValue(pricing[i].NormalValue, "BRL"),
			UsdValue: qr.convertValue(pricing[i].NormalValue, "USD"),
			EurValue: qr.convertValue(pricing[i].NormalValue, "EUR"),
		}

		professional := &structs.Charge{
			Name:     pricing[i].Title,
			BrlValue: qr.convertValue(pricing[i].PremiumValue, "BRL"),
			UsdValue: qr.convertValue(pricing[i].PremiumValue, "USD"),
			EurValue: qr.convertValue(pricing[i].PremiumValue, "EUR"),
		}

		basicPlanCharges[i] = basic
		professionalPlanCharges[i] = professional
	}

	siteData := &structs.Site{
		URL: *site,
		Plans: []*structs.Plans{
			{Name: "Básico", Charges: basicPlanCharges},
			{Name: "Profissinal", Charges: professionalPlanCharges},
		},
	}

	found := []*structs.Site{siteData}
	return found, nil
}

func (qr *QueryResolver) convertValue(value string, currency string) string {
	if value == "--" {
		return "Não disponvel para este plano"
	}

	i, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return "Consulte-nos"
	}

	priceTag := qr.monetary.Convert(i, currency)
	return fmt.Sprintf("%s: %.2f", currencySigns[currency], priceTag)
}
