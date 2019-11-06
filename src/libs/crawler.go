package libs

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Charge representation
type Charge struct {
	Title        string
	NormalValue  string
	PremiumValue string
}

// GetPricing returns product pricing
func GetPricing(url string) {
	var charges []Charge

	data := readBody(url)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		fmt.Println("No url found")
		log.Fatal(err)
	}

	doc.Find("#tarifas-2").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find(".row").Each(func(index int, tablerow *goquery.Selection) {
			if index > 0 {
				chargeData := Charge{}
				tablerow.Find(".col-sm-4").Each(func(index int, data *goquery.Selection) {
					switch index {
					case 0:
						chargeData.Title = data.Text()
					case 1:
						chargeData.NormalValue = getValues(data.Text())
					case 2:
						chargeData.PremiumValue = getValues(data.Text())
					}
				})

				charges = append(charges, chargeData)
			}
		})
	})
}

func getValues(html string) string {
	var rgx = regexp.MustCompile(`(\d+((\.|,|)(\d*|)))`)
	rs := rgx.FindStringSubmatch(html)
	if len(rs) < 1 {
		return strings.TrimSpace(html)
	}

	return strings.Replace(rs[1], ",", ".", -1)
}

func readBody(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("ERROR: ", err)
		return ""
	}

	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR: ", err)
		return ""
	}

	return string(html)
}
