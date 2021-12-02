package plivo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBrand_List(t *testing.T) {
	expectResponse("brandListResponse.json", 200)
	assert := require.New(t)
	resp, err := client.Brand.List(BrandListParams{})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.BrandResponse[0].BrandID)
	cl := client.httpClient
	client.httpClient = nil
	brandType := "PRIVATE_PROFIT"
	status := "VERIFIED"
	resp, err = client.Brand.List(BrandListParams{Type: &brandType, Status: &status})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "GET", "10dlc/Brand")
}

func TestBrand_Get(t *testing.T) {
	expectResponse("brandGetResponse.json", 200)
	BrandID := "BRPXS6E"
	assert := require.New(t)
	brand, err := client.Brand.Get(BrandID)
	assert.NotNil(brand)
	assert.Nil(err)
	assert.Equal(BrandID, brand.Brand.BrandID)
	assert.NotEmpty(brand.ApiID)

	cl := client.httpClient
	client.httpClient = nil
	brand, err = client.Brand.Get(BrandID)
	assert.NotNil(err)
	assert.Nil(brand)
	client.httpClient = cl

	assertRequest(t, "GET", "10dlc/Brand/%s", BrandID)
}

func TestBrand_Create(t *testing.T) {
	expectResponse("brandCreationResponse.json", 200)
	assert := require.New(t)
	altBusinessIDType := "GIIN"
	altBusinessID := "111"
	firstName := "John"
	lastName := "Doe"
	website := "http://www.abcmobile.com"
	resp, err := client.Brand.Create(BrandCreationParams{
		AltBusinessIDType:  &altBusinessIDType,
		AltBusinessID:      &altBusinessID,
		City:               "New York",
		CompanyName:        "ABC Inc.",
		Country:            "US",
		Ein:                "111111111",
		EinIssuingCountry:  "US",
		Email:              "johndoe@abc.com",
		EntityType:         "PRIVATE_PROFIT",
		FirstName:          &firstName,
		LastName:           &lastName,
		Phone:              "+11234567890",
		PostalCode:         "10001",
		RegistrationStatus: "PENDING",
		State:              "NY",
		StockExchange:      "NASDAQ",
		StockSymbol:        "ABC",
		Street:             "123",
		Vertical:           "RETAIL",
		Website:            &website,
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	assert.NotEmpty(resp.Brand)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Brand.Create(BrandCreationParams{
		AltBusinessIDType:  &altBusinessIDType,
		AltBusinessID:      &altBusinessID,
		City:               "New York",
		CompanyName:        "ABC Inc.",
		Country:            "US",
		Ein:                "111111111",
		EinIssuingCountry:  "US",
		Email:              "johndoe@abc.com",
		EntityType:         "PRIVATE_PROFIT",
		FirstName:          &firstName,
		LastName:           &lastName,
		Phone:              "+11234567890",
		PostalCode:         "10001",
		RegistrationStatus: "PENDING",
		State:              "NY",
		StockExchange:      "NASDAQ",
		StockSymbol:        "ABC",
		Street:             "123",
		Vertical:           "RETAIL",
		Website:            &website,
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "10dlc/Brand")
}
