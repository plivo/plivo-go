package plivo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProfile_List(t *testing.T) {
	expectResponse("profileListResponse.json", 200)
	assert := require.New(t)
	resp, err := client.Profile.List(ProfileListParams{Limit: 2, Offset: 0})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ProfileResponse[0].ProfileUUID)
	cl := client.httpClient
	client.httpClient = nil

	resp, err = client.Profile.List(ProfileListParams{Limit: 2, Offset: 0})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "GET", "Profile")
}

func TestProfile_Get(t *testing.T) {
	expectResponse("profileGetResponse.json", 200)
	ProfileUUID := "201faedc-7df9-4840-9ab1-3997ce3f7cf4"
	assert := require.New(t)
	profile, err := client.Profile.Get(ProfileUUID)
	assert.NotNil(profile)
	assert.Nil(err)
	assert.Equal(ProfileUUID, profile.Profile.ProfileUUID)
	assert.NotEmpty(profile.ApiID)

	cl := client.httpClient
	client.httpClient = nil
	profile, err = client.Profile.Get(ProfileUUID)
	assert.NotNil(err)
	assert.Nil(profile)
	client.httpClient = cl

	assertRequest(t, "GET", "Profile/%s", ProfileUUID)
}

func TestProfile_Create(t *testing.T) {
	expectResponse("profileCreationResponse.json", 200)
	assert := require.New(t)
	resp, err := client.Profile.Create(CreateProfileRequestParams{
		ProfileAlias:      "used_in_automation_01",
		CustomerType:      "DIRECT",
		EntityType:        "PRIVATE",
		CompanyName:       "ABC Inc.",
		Vertical:          "ENERGY",
		StockSymbol:       "ABC",
		StockExchange:     "NSE",
		AltBusinessidType: "NONE",
		Ein:               "123456789",
		EinIssuingCountry: "US",
		Website:           "www.google.com",
		Address: &Address{
			Street:     "123",
			City:       "New York",
			State:      "NY",
			PostalCode: "10001",
			Country:    "IN",
		},
		AuthorizedContact: &AuthorizedContact{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "test@plivo.com",
			Title:     "Doe",
			Seniority: "admin",
			Phone:     "919539113734",
		},
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	assert.NotEmpty(resp.ProfileUUID)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Profile.Create(CreateProfileRequestParams{
		ProfileAlias:      "used_in_automation_01",
		CustomerType:      "DIRECT",
		EntityType:        "PRIVATE",
		CompanyName:       "ABC Inc.",
		Vertical:          "ENERGY",
		StockSymbol:       "ABC",
		StockExchange:     "NSE",
		AltBusinessidType: "NONE",
		Ein:               "123456789",
		EinIssuingCountry: "US",
		Website:           "www.google.com",
		Address: &Address{
			Street:     "123",
			City:       "New York",
			State:      "NY",
			PostalCode: "10001",
			Country:    "IN",
		},
		AuthorizedContact: &AuthorizedContact{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "test@plivo.com",
			Title:     "Doe",
			Seniority: "admin",
			Phone:     "919539113734",
		},
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "Profile")
}

func TestProfile_Delete(t *testing.T) {
	expectResponse("profileDeleteResponse.json", 200)
	assert := require.New(t)
	ProfileUUID := "201faedc-7df9-4840-9ab1-3997ce3f7cf4"
	resp, err := client.Profile.Delete(ProfileUUID)
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Profile.Delete(ProfileUUID)
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "DELETE", "Profile/201faedc-7df9-4840-9ab1-3997ce3f7cf4")

}

func TestProfile_Update(t *testing.T) {
	expectResponse("profileUpdateResponse.json", 200)
	assert := require.New(t)
	resp, err := client.Profile.Update("09322f43-fe16-4525-b8e4-4229c867795d", UpdateProfileRequestParams{
		EntityType:  "PRIVATE",
		CompanyName: "ABC Inc.",
		Vertical:    "ENERGY",
		Website:     "www.google.com",
		Address: &Address{
			Street:     "123",
			City:       "New York",
			State:      "NY",
			PostalCode: "10001",
			Country:    "IN",
		},
		AuthorizedContact: &AuthorizedContact{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "test@plivo.com",
			Title:     "Doe",
			Seniority: "admin",
			Phone:     "919539113734",
		},
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Profile.Update("09322f43-fe16-4525-b8e4-4229c867795d", UpdateProfileRequestParams{
		EntityType:  "PRIVATE",
		CompanyName: "ABC Inc.",
		Vertical:    "ENERGY",
		Website:     "www.google.com",
		Address: &Address{
			Street:     "123",
			City:       "New York",
			State:      "NY",
			PostalCode: "10001",
			Country:    "IN",
		},
		AuthorizedContact: &AuthorizedContact{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "test@plivo.com",
			Title:     "Doe",
			Seniority: "admin",
			Phone:     "919539113734",
		},
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "Profile")
}
