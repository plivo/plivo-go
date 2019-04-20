package main

import (
	"encoding/json"
	"fmt"
	"github.com/plivo/plivo-go"
)

// Initialize the AuthId and AuthToken parameters in credentials.go
// To trigger Application resource methods invoke corresponding helper method in main.go


// Example to create a Application
// Pass ApplicationCreateParams to create a application
// Application name and answer url are mandate params

func TestCreateApplication(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Applications.Create(
		plivo.ApplicationCreateParams{
			AppName: "Test Application suhas Test",
			AnswerURL: "https://answer.url",
			LogIncomingMessages:false,
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to get a Application
//Pass ApplicationID of corresponding application

func TestApplicationGet(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Applications.Get(
		"25387659147377995",
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)

}

//Example to update a Application
//Pass ApplicationUpdateParams and corresponding ApplicationID  to update a application

func TestApplicationUpdate(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Applications.Update(
		"25387659147377995",
		plivo.ApplicationUpdateParams{
			LogIncomingMessages:true,
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to delete a Application
//Pass corresponding ApplicationID to delete a application

func TestApplicationDelete(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	err = client.Applications.Delete(
		"app_id",
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Application deleted successfully.")
}

//Example to list all Applications
//Pass offset and limit values to limit the number of applications

func TestApplicationList(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Applications.List(
		plivo.ApplicationListParams{
			Offset: 0,
			Limit: 5,
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}