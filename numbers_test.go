package plivo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberService_Create(t *testing.T) {
	expectResponse("NumberCreateResponse.json", 202)

	if _, err := client.Numbers.Create(NumberCreateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Numbers.Create(NumberCreateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Number")
}

func TestNumberService_Update(t *testing.T) {
	expectResponse("NumberUpdateResponse.json", 202)
	NumberId := "NumberId"

	if _, err := client.Numbers.Update(NumberId, NumberUpdateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Numbers.Update(NumberId, NumberUpdateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Number/%s", NumberId)
}

func TestNumberService_List(t *testing.T) {
	expectResponse("NumberListResponse.json", 200)

	if _, err := client.Numbers.List(NumberListParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Numbers.List(NumberListParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Number")
}

func TestNumberService_Get(t *testing.T) {
	expectResponse("NumberGetResponse.json", 200)
	NumberId := "NumberId"

	number, err := client.Numbers.Get(NumberId)
	assert.Equal(t, number.ID(), number.Number)
	if err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err = client.Numbers.Get(NumberId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Number/%s", NumberId)
}

func TestNumberService_Delete(t *testing.T) {
	expectResponse("", 204)
	NumberId := "NumberId"

	if err := client.Numbers.Delete(NumberId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	err := client.Numbers.Delete(NumberId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Number/%s", NumberId)
}

func TestPhoneNumberService_Create(t *testing.T) {
	expectResponse("PhoneNumberCreateResponse.json", 202)

	if _, err := client.PhoneNumbers.Create("123", PhoneNumberCreateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.PhoneNumbers.Create("123", PhoneNumberCreateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "PhoneNumber/%s", "123")
}

func TestPhoneNumberService_List(t *testing.T) {
	expectResponse("PhoneNumberListResponse.json", 200)

	if _, err := client.PhoneNumbers.List(PhoneNumberListParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.PhoneNumbers.List(PhoneNumberListParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "PhoneNumber")
}
