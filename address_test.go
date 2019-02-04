package plivo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressService_Create(t *testing.T) {
	expectResponse("addressCreateResponse.json", 202)

	if _, err := client.Addresses.Create(AddressCreateParams{
		PhoneNumberCountry: "test",
		NumberType:         AddressNumberType.LOCAL,
		Salutation:         AddressSalutationType.MR,
		AddressLine1:       "test",
		FirstName:          "test",
		LastName:           "test",
		AddressLine2:       "test",
		City:               "test",
		Region:             "test",
		PostalCode:         "test",
		CountryIso:         "IN",
	},); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Addresses.Create(AddressCreateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Verification/Address")
}

func TestAddressService_Get(t *testing.T) {
	expectResponse("addressGetResponse.json", 202)
	documentID := "documentID"

	address, err := client.Addresses.Get(documentID)
	assert.Equal(t, address.DocumentID(), address.ID)
	if err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err = client.Addresses.Get(documentID)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Verification/Address/%s", documentID)
}

func TestAddressService_Update(t *testing.T) {
	expectResponse("addressModifyResponse.json", 202)
	documentID := "documentID"

	if _, err := client.Addresses.Update(documentID, AddressUpdateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Addresses.Update(documentID, AddressUpdateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Verification/Address/%s", documentID)
}

func TestAddressService_Delete(t *testing.T) {
	expectResponse("", 204)
	documentID := "documentID"

	if err := client.Addresses.Delete(documentID); err != nil {
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

	assertRequest(t, "DELETE", "Verification/Address/%s", documentID)
}

func TestAddressService_List(t *testing.T) {
	expectResponse("addressListResponse.json", 200)

	if _, err := client.Addresses.List(AddressListParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Addresses.List(AddressListParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Verification/Address/")
}
