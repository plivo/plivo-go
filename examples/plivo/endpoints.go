package main

import (
	"encoding/json"
	"fmt"
	"github.com/plivo/plivo-go"
)

// Initialize the AuthId and AuthToken parameters in credentials.go
// To trigger Conference resource methods invoke corresponding helper method in main()
// Initialize the following the endPointId to trigger appropriate helper methods

// 	To build and run endpoints.go
//  cd  plivo-go/examples/plivo
//  go run endpoints.go credentials.go

const endPointId = "373XXXXX03666"

func main(){
	testEndpointGet()
}

//Example to create a endpoint
//Username , Password, Alias are mandatory params
func testEndpointCreate(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Endpoints.Create(
		plivo.EndpointCreateParams{
			Username: "testusername",
			Password: "testpassword",
			Alias: "Test Account",
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to get details a endpoint
//Pass corresponding endpointID to fetch endpoint details
func testEndpointGet(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Endpoints.Get(
		endPointId,
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))

}

//Example to update a endpoint
//Pass corresponding endpointID to update
func testEndpointUpdate(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Endpoints.Update(
		endPointId,
		plivo.EndpointUpdateParams{
			Alias: "Updated Endpoint Alias",
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to list all  endpoints

func testEndpointList(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Endpoints.List(
		plivo.EndpointListParams{},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to delete a endpoint
//Pass corresponding endpointID to delete
func testEndpointDelete(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	err = client.Endpoints.Delete(
		endPointId,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted successfully.")
}