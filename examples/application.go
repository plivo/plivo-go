package main

import (
	"fmt"
	"os"

	"github.com/plivo/plivo-go"
)

var client *plivo.Client

func initClient(authID, authToken string) {
	var er error
	copts := &plivo.ClientOptions{}
	client, er = plivo.NewClient(authID, authToken, copts)
	if er != nil {
		panic(er)
	}
}
func CreateApplication(params plivo.ApplicationCreateParams) (*plivo.ApplicationCreateResponseBody, error) {
	response, er := client.Applications.Create(params)
	if er != nil {
		return nil, er
	}
	return response, nil
}

func UpdateApplication(appId string, params plivo.ApplicationUpdateParams) (*plivo.ApplicationUpdateResponse, error) {
	response, er := client.Applications.Update(appId, params)
	if er != nil {
		return nil, er
	}
	return response, nil
}

func GetApplication(appId string) (*plivo.Application, error) {
	response, er := client.Applications.Get(appId)
	if er != nil {
		return nil, er
	}
	return response, nil
}

func applicationMain() {
	initClient("", "")
	input := plivo.ApplicationCreateParams{
		AppName:   "example-app",
		AnswerURL: "https://www.my-answer-url.com/",
	}

	// create app
	createResponse, er := CreateApplication(input)
	if er != nil {
		fmt.Printf("Error occurred while creating application. error:%+v\n", er)
		os.Exit(1)
	} else {
		fmt.Printf("%+v\n", createResponse)
	}

	// get app
	appDetails, er := GetApplication(createResponse.AppID)
	if er != nil {
		fmt.Printf("Error occurred while getting application details. error:%+v\n", er)
		os.Exit(1)
	} else {
		fmt.Printf("%+v\n", appDetails)
	}

	// update app
	updatedApp := plivo.ApplicationUpdateParams{
		DefaultNumberApp: true,
		AppName:          "default-example-app",
		AnswerURL:        "https://www.my-answer-url.com/default_app/",
	}
	updateResponse, er := UpdateApplication(createResponse.AppID, updatedApp)
	if er != nil {
		fmt.Printf("Error occurred while updating application. error:%+v\n", er)
		os.Exit(1)
	} else {
		fmt.Printf("%+v\n", updateResponse)
	}

	// make an app public
	updatedApp.PublicURI = true
	updateResponse, er = UpdateApplication(createResponse.AppID, updatedApp)
	if er != nil {
		fmt.Printf("Error occurred while making the application public. error:%+v\n", er)
		os.Exit(1)
	} else {
		fmt.Printf("%+v\n", updateResponse)
	}

	appDetails, er = GetApplication(createResponse.AppID)
	if er != nil {
		fmt.Printf("Error occurred while getting application details. error:%+v\n", er)
		os.Exit(1)
	} else {
		fmt.Printf("%+v\n", appDetails)
	}
}
