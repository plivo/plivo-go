package main

import (
	"fmt"
	"encoding/json"
	"github.com/plivo/plivo-go"
)

// Initialize the AuthId and AuthToken parameters in credentials.go
// To trigger Account resource methods invoke corresponding helper method in main()


// 	To build and run accounts.go
//  cd  plivo-go/examples/plivo
//  go run accounts.go credentials.go


func main() {
	testAccountGet()
}

//Example to get a Account details

func testAccountGet(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Accounts.Get()
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println(string(responseJSON))
}

//Example to update Account details
//To update use AccountUpdateParams

func testAccountupdate(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Accounts.Update(
		plivo.AccountUpdateParams{
			Name: "Lucius Fox",
			City: "New York",
			Address: "Times Square",
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println(string(responseJSON))
}

//Example to get Sub-Account details
//To get sub-account details pass sub_account_id

func testSubAccountGet(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response,err := client.Subaccounts.Get("SAZDIXMTUXXXXXHYJFJNG")
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println(string(responseJSON))
}

//Example to create sub-account
//Create a SubAccount with SubaccountCreateParams

func testCreateSubAccount(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Subaccounts.Create(
		plivo.SubaccountCreateParams{
			Name: "Test Subaccount",
			Enabled: true,
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println(string(responseJSON))
}

//Example to update a sub-ccount
//Pass sub-account-id and SubaccountUpdateParams to be updated

func testUpdateSubAccount(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Subaccounts.Update(
		"SAZDINXXXXXXXXFJNG",
		plivo.SubaccountUpdateParams{
			Name: "sub_account_name",
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}

//Example to delete a sub-account-id
// Pass corresponding sub-account-id to be deleted

func testDeleteSubAccount(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	err = client.Subaccounts.Delete("SAZDINXXXXXXXXFJNG")
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted SubAccount successfully.")
}

//Example to list sub accounts
//Pass offset and limit values to limit the number of sub-accounts

func testListAllSubAccount(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Subaccounts.List(
		plivo.SubaccountListParams{
			Offset: 0,
			Limit: 5,
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println(string(responseJSON))
}