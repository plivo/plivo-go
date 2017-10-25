package plivo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPricingService_Get(t *testing.T) {
	expectResponse("PricingGetResponse.json", 200)

	pricing, err := client.Pricing.Get("GB")
	assert.Equal(t, pricing.ID(), pricing.CountryISO)
	if err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err = client.Pricing.Get("GB")
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Pricing")
}
