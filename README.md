# plivo-go

[![Build, Unit Tests, Linters Status](https://github.com/plivo/plivo-go/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/plivo/plivo-go/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/plivo/plivo-go/branch/master/graph/badge.svg)](https://codecov.io/gh/plivo/plivo-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/plivo/plivo-go)](https://goreportcard.com/report/github.com/plivo/plivo-go)
[![GoDoc](https://godoc.org/github.com/plivo/plivo-go?status.svg)](https://godoc.org/github.com/plivo/plivo-go)

The Plivo Go SDK makes it simpler to integrate communications into your Go applications using the Plivo REST API. Using the SDK, you will be able to make voice calls, send messages and generate Plivo XML to control your call flows.

## Prerequisites

- Go >= 1.13.x

## Getting started

The steps described below uses go modules.

###### Create a new project (optional)
```sh
$ mkdir ~/helloplivo
$ cd ~/helloplivo
$ go mod init helloplivo
```

This will generate a `go.mod` and `go.sum` file.

###### Add plivo-go as a dependency to your project
```sh
$ go get github.com/plivo/plivo-go/v7
```

### Authentication
To make the API requests, you need to create a `Client` and provide it with authentication credentials (which can be found at [https://manage.plivo.com/dashboard/](https://manage.plivo.com/dashboard/)).

We recommend that you store your credentials in the `PLIVO_AUTH_ID` and the `PLIVO_AUTH_TOKEN` environment variables, so as to avoid the possibility of accidentally committing them to source control. If you do this, you can initialise the client with no arguments and it will automatically fetch them from the environment variables:

```go
package main

import "github.com/plivo/plivo-go/v7"

func main()  {
	client, err := plivo.NewClient("", "", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
}
```
Alternatively, you can specifiy the authentication credentials while initializing the `Client`.

```go
package main

import "github.com/plivo/plivo-go/v7"

func main()  {
	client, err := plivo.NewClient("<auth-id>", "<auth-token>", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
}
```

## The Basics
The SDK uses consistent interfaces to create, retrieve, update, delete and list resources. The pattern followed is as follows:

```go
client.Resources.Create(Params{}) // Create
client.Resources.Get(Id) // Get
client.Resources.Update(Id, Params{}) // Update
client.Resources.Delete(Id) // Delete
client.Resources.List() // List all resources, max 20 at a time
```

Using `client.Resources.List()` would list the first 20 resources by default (which is the first page, with `limit` as 20, and `offset` as 0). To get more, you will have to use `limit` and `offset` to get the second page of resources.

## Examples

### Send a message

```go
package main

import "github.com/plivo/plivo-go/v7"

func main() {
	client, err := plivo.NewClient("", "", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	client.Messages.Create(plivo.MessageCreateParams{
		Src:  "the_source_number",
		Dst:  "the_destination_number",
		Text: "Hello, world!",
	})
}
```

### Make a call

```go
package main

import "github.com/plivo/plivo-go/v7"

func main() {
	client, err := plivo.NewClient("", "", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	client.Calls.Create(plivo.CallCreateParams{
		From:      "the_source_number",
		To:        "the_destination_number",
		AnswerURL: "http://answer.url",
	})
}

```

### Lookup a number

```go
package main

import (
	"fmt"
	"log"

	"github.com/plivo/plivo-go/v7"
)

func main() {
	client, err := plivo.NewClient("<auth-id>", "<auth-token>", &plivo.ClientOptions{})
	if err != nil {
		log.Fatalf("plivo.NewClient() failed: %s", err.Error())
	}

	resp, err := client.Lookup.Get("<insert-number-here>", plivo.LookupParams{})
	if err != nil {
		if respErr, ok := err.(*plivo.LookupError); ok {
			fmt.Printf("API ID: %s\nError Code: %d\nMessage: %s\n",
				respErr.ApiID, respErr.ErrorCode, respErr.Message)
			return
		}
		log.Fatalf("client.Lookup.Get() failed: %s", err.Error())
	}

	fmt.Printf("%+v\n", resp)
}
```

### Generate Plivo XML

```go
package main

import "github.com/plivo/plivo-go/v7/xml"

func main() {
	println(xml.ResponseElement{
		Contents: []interface{}{
			new(xml.SpeakElement).SetContents("Hello, world!"),
		},
	}.String())
}
```

This generates the following XML:

```xml
<Response>
  <Speak>Hello, world!</Speak>
</Response>
```

### Run a PHLO

```go
package main

import (
	"fmt"
	"github.com/plivo/plivo-go/v7"
)

// Initialize the following params with corresponding values to trigger resources

const authId = "auth_id"
const authToken = "auth_token"
const phloId = "phlo_id"

// with payload in request

func main() {
	testPhloRunWithParams()
}

func testPhloRunWithParams() {
	phloClient, err := plivo.NewPhloClient(authId, authToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get(phloId)
	if err != nil {
		panic(err)
	}
	//pass corresponding from and to values
	type params map[string]interface{}
	response, err := phloGet.Run(params{
		"from": "111111111",
		"to":   "2222222222",
	})

	if err != nil {
		println(err)
	}
	fmt.Printf("Response: %#v\n", response)
}
```

### More examples
Refer to the [Plivo API Reference](https://www.plivo.com/docs/sms/api/overview/) for more examples.

## Local Development
> Note: Requires latest versions of Docker & Docker-Compose. If you're on MacOS, ensure Docker Desktop is running.
1. Export the following environment variables in your host machine:
```bash
export PLIVO_AUTH_ID=<your_auth_id>
export PLIVO_AUTH_TOKEN=<your_auth_token>
export PLIVO_API_DEV_HOST=<plivoapi_dev_endpoint>
export PLIVO_API_PROD_HOST=<plivoapi_public_endpoint>
```
2. Run `make build`. This will create a docker container in which the sdk will be setup and dependencies will be installed.
> The entrypoint of the docker container will be the `setup_sdk.sh` script. The script will handle all the necessary changes required for local development.
3. The above command will print the docker container id (and instructions to connect to it) to stdout.
4. The testing code can be added to `<sdk_dir_path>/go-sdk-test/test.go` in host  
 (or `/usr/src/app/go-sdk-test/test.go` in container)
5. The sdk directory will be mounted as a volume in the container. So any changes in the sdk code will also be reflected inside the container.
6. To run unit tests, run `make test CONTAINER=<cont_id>` in host, where `<cont_id>` is the docker container id created in 2.   
(The docker container should be running)