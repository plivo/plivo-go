package plivo

import (
"errors"
"github.com/stretchr/testify/assert"
"testing"
)

func TestIdentityService_Create(t *testing.T) {
	expectResponse("identityCreateResponse.json", 202)

	if _, err := client.Identities.Create(IdentityCreateParams{
		PhoneNumberCountry: "test",
		NumberType:         IdentityNumberType.LOCAL,
		Salutation:         IdentitySalutationType.MR,
		AddressLine1:       "test",
		FirstName:          "test",
		LastName:           "test",
		AddressLine2:       "test",
		City:               "test",
		Region:             "test",
		PostalCode:         "test",
		CountryIso:         "IN",
		AddressProofType:   IdentityProofType.PASSPORT,
		IdNumber:           "346234",
		Nationality:        "IN",
	},); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Identities.Create(IdentityCreateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Verification/Identity")
}

func TestIdentityService_Get(t *testing.T) {
	expectResponse("identityGetResponse.json", 202)
	documentID := "documentID"

	identity, err := client.Identities.Get(documentID)
	assert.Equal(t, identity.DocumentID(), identity.ID)
	if err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err = client.Identities.Get(documentID)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Verification/Identity/%s", documentID)
}

func TestIdentityService_Update(t *testing.T) {
	expectResponse("identityModifyResponse.json", 202)
	documentID := "documentID"

	if _, err := client.Identities.Update(documentID, IdentityUpdateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Identities.Update(documentID, IdentityUpdateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Verification/Identity/%s", documentID)
}

func TestIdentityService_Delete(t *testing.T) {
	expectResponse("", 204)
	documentID := "documentID"

	if err := client.Identities.Delete(documentID); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	err := client.Addresses.Delete(documentID)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Verification/Identity/%s", documentID)
}

func TestIdentityService_List(t *testing.T) {
	expectResponse("identityListResponse.json", 200)

	if _, err := client.Identities.List(IdentityListParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Identities.List(IdentityListParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Verification/Identity/")
}
