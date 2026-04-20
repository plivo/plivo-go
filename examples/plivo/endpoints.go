package main

import (
	"encoding/json"
	"fmt"
	"github.com/plivo/plivo-go"
)

// Initialize the AuthId and AuthToken parameters in credentials.go
// To trigger Conference resource methods invoke corresponding helper method in main.go
// Initialize the following the endPointId to trigger appropriate helper methods

const endPointId = "373XXXXX03666"

//Example to create a endpoint
//Username , Password, Alias are mandatory params

func TestEndpointCreate(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Endpoints.Create(
		plivo.EndpointCreateParams{
			Username: "Testusername",
			Password: "Testpassword",
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
func TestEndpointGet(){
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
func TestEndpointUpdate(){
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

func TestEndpointList(){
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
func TestEndpointDelete(){
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