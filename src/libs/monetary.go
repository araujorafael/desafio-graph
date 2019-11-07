package libs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const apiLink = "https://api.exchangeratesapi.io/latest?base=BRL"

// Monetary interface
type Monetary interface {
	Convert(value float64, currency string) float64
}

// MonetaryImpl struct implementation
type MonetaryImpl struct{}

type jsonResp struct {
	Rates map[string]float64 `json:"rates"`
}

// Convert givern currency value
func (m MonetaryImpl) Convert(value float64, currency string) float64 {
	rates := m.getRates()
	return value / rates[currency]
}

func (m MonetaryImpl) getRates() map[string]float64 {
	response, err := http.Get(apiLink)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonData := jsonResp{}
	err = json.Unmarshal(resp, &jsonData)
	if err != nil {
		log.Fatal(err)
	}

	return jsonData.Rates
}
