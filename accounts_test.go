package plivo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountService_Get(t *testing.T) {
	expectResponse("accountGetResponse.json", 200)

	account, err := client.Accounts.Get()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, account.ID(), account.AuthID)

	cl := client.httpClient
	client.httpClient = nil
	_, err = client.Accounts.Get()
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "")
}

func TestAccountService_Update(t *testing.T) {
	expectResponse("accountModifyResponse.json", 200)

	if _, err := client.Accounts.Update(AccountUpdateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Accounts.Update(AccountUpdateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "")
}

func TestSubaccountService_Get(t *testing.T) {
	expectResponse("subaccountGetResponse.json", 202)
	subauthId := "subauthId"

	subaccount, err := client.Subaccounts.Get(subauthId)
	assert.Equal(t, subaccount.ID(), subaccount.AuthID)
	if err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err = client.Subaccounts.Get(subauthId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Subaccount/%s", subauthId)
}

func TestSubaccountService_List(t *testing.T) {
	expectResponse("subaccountListResponse.json", 202)

	if _, err := client.Subaccounts.List(SubaccountListParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Subaccounts.List(SubaccountListParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Subaccount")
}

func TestSubaccountService_Create(t *testing.T) {
	expectResponse("subaccountCreateResponse.json", 202)

	if _, err := client.Subaccounts.Create(SubaccountCreateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Subaccounts.Create(SubaccountCreateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Subaccount")
}

func TestSubaccountService_Update(t *testing.T) {
	expectResponse("subaccountModifyResponse.json", 202)
	subauthId := "subauthId"

	if _, err := client.Subaccounts.Update(subauthId, SubaccountUpdateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Subaccounts.Update(subauthId, SubaccountUpdateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Subaccount/%s", subauthId)
}

func TestSubaccountService_Delete(t *testing.T) {
	expectResponse("", 202)
	subauthId := "subauthId"

	if err := client.Subaccounts.Delete(subauthId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	err := client.Subaccounts.Delete(subauthId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Subaccount/%s", subauthId)
}
