package resolvers_test

import (
	"context"
	"desafio-graph/mocks"
	"desafio-graph/src/libs"
	"desafio-graph/src/resolvers"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

var crawler *mocks.Crawler
var ctx context.Context

func TestSitesErr(t *testing.T) {
	var (
		sitesResolver *resolvers.QueryResolver
		monetary      *mocks.Monetary
		crawler       *mocks.Crawler
	)

	t.Run("Should call crawler and monetary libs lib", func(t *testing.T) {
		pricing := []libs.Charge{
			{Title: "Nota", NormalValue: "5.00", PremiumValue: "6.00"},
			{Title: "Nota", NormalValue: "5.00", PremiumValue: "6.00"},
		}
		crawler := new(mocks.Crawler)
		crawler.On("GetPricing", mock.Anything).Return(pricing)

		monetary := new(mocks.Monetary)
		monetary.On("Convert", mock.Anything, mock.Anything).Return(89.89)

		sitesResolver = resolvers.NewQueryResolver(crawler, monetary)
		site := "https://www.smartmei.com.br"
		_, err := sitesResolver.Sites(ctx, &site)

		if err != nil {
			t.Error() // to indicate test failed
		} else {
			t.Logf("Successfully returned error when site is not found")
		}

		crawler.AssertCalled(t, "GetPricing", site)
		fmt.Printf("\n\n %#v \n\n", monetary)
		monetary.AssertNumberOfCalls(t, "Convert", 12)
	})

	t.Run("Should return error when site is not given", func(t *testing.T) {
		sitesResolver = resolvers.NewQueryResolver(crawler, monetary)
		site := ""
		_, err := sitesResolver.Sites(ctx, &site)

		if err == nil {
			t.Error("Site not given, should return a error")
		} else {
			t.Logf("Successfully returned error when site is empty")
		}
	})

	t.Run("Should return error when site is not found", func(t *testing.T) {
		sitesResolver = resolvers.NewQueryResolver(crawler, monetary)
		site := "https://notfound.com"
		_, err := sitesResolver.Sites(ctx, &site)

		if err == nil {
			t.Error() // to indicate test failed
		} else {
			t.Logf("Successfully returned error when site is not found")
		}
	})
}
