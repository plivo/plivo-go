# plivo-go

[![Build Status](https://travis-ci.org/plivo/plivo-go.svg?branch=master)](https://travis-ci.org/plivo/plivo-go)
[![GoDoc](https://godoc.org/github.com/plivo/plivo-go?status.svg)](https://godoc.org/github.com/plivo/plivo-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/plivo/plivo-go)](https://goreportcard.com/report/github.com/plivo/plivo-go)

The Plivo Go SDK makes it simpler to integrate communications into your Go applications using the Plivo REST API. Using the SDK, you will be able to make voice calls, send SMS and generate Plivo XML to control your call flows.

## Prerequisites

- Go >= 1.7.x

## Installation

### To Install Stable release

You can use the Stable release using the `go` command.

    $ go get github.com/plivo/plivo-go

You can also install by cloning this repository into your `GOPATH`.

### To Install Beta release

1. In terminal, using the following command, create a new folder called **test-plivo-beta**.

    ```
    $ mkdir test-plivo-beta
    ```

	**Note:** Make sure the new folder is outside your GOPATH.

2. Change your directory to the new folder.
3. Using the following command, initialize a new module:

    ```
    $ go mod init github.com/plivo/beta
    ```

    You will see the following return:

    ```
    $ go mod init github.com/plivo/beta
    ```

4. Next, create a new go file with the following code:

    ```go
    package main

    import (
    "fmt"
    "github.com/plivo/plivo-go"
    )

    const authId  = "auth_id"
    const authToken = "auth_token"

    func main() {
    	testPhloGetter()
    }

    func testPhloGetter() {
    	phloClient,err := plivo.NewPhloClient(authId, authToken, &plivo.ClientOptions{})
    	if err != nil {
    		panic(err)
    	}
    	response, err := phloClient.Phlos.Get("phlo_id")
    	if err != nil {
    		panic(err)
    	}
    	fmt.Printf("Response: %#v\n", response)
    }
    ```

    **Note:** Make sure you enter your `auth_id` and `auth_token` before you initialize the code.

5. Run the following command to build the packages:

    ```
    $ go build
    ```
    <img id="myImg" src="https://s3.amazonaws.com/plivo_blog_uploads/static_assets/images/server_sdks/step5.png" alt="payload_defined" style="border: 1px solid #e8eaf1;">

    A file named go.mod will be generated.

6. Open go.mod using the command `vim go.mod` and edit the **plivo-go** version to the beta version you want to download.

    <img id="myImg" src="https://s3.amazonaws.com/plivo_blog_uploads/static_assets/images/server_sdks/step6.png" alt="payload_defined" style="border: 1px solid #e8eaf1;">

    **Note:** You can see the full list of releases [here](https://github.com/plivo/plivo-go/releases).

    For example, change

    ```
    github.com/plivo/plivo-go v4.0.5+incompatible 
    ```

    to 

    ```
    github.com/plivo/plivo-go v4.0.6-beta1
    ```

7. Once done, save the go.mod.
8. Run **go build** to build the packages.

    **go.mod** will be updated with the beta version.

You can now use the features available in the Beta branch.

## Getting started

### Authentication
To make the API requests, you need to create a `Client` and provide it with authentication credentials (which can be found at [https://console.plivo.com/dashboard/](https://console.plivo.com/dashboard/)).

We recommend that you store your credentials in the `PLIVO_AUTH_ID` and the `PLIVO_AUTH_TOKEN` environment variables, so as to avoid the possibility of accidentally committing them to source control. If you do this, you can initialise the client with no arguments and it will automatically fetch them from the environment variables:

```go
package main

import "github.com/plivo/plivo-go"

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

import "github.com/plivo/plivo-go"

func main()  {
 client, err := plivo.NewClient("<auth_id>", "<auth_token>", &plivo.ClientOptions{})
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

import "github.com/plivo/plivo-go"

func main()  {
  client, err := plivo.NewClient("<auth_id>", "<auth_token>", &plivo.ClientOptions{})
  if err != nil {
    panic(err)
  }
  client.Messages.Create(plivo.MessageCreateParams{
    Src: "+14156667778",
    Dst: "+14156667777",
    Text: "Hello, world!",
  })
}
```

### Make a call

```go
package main

import "github.com/plivo/plivo-go"

func main()  {
  client, err := plivo.NewClient("<auth_id>", "<auth_token>", &plivo.ClientOptions{})
  if err != nil {
    panic(err)
  }
  client.Calls.Create(plivo.CallCreateParams{
    From: "+14156667778",
    To: "+14156667777",
    AnswerURL: "http://s3.amazonaws.com/static.plivo.com/answer.xml",
    Method: "GET"
  })
}
```

### Lookup a number

```go
package main

import (
	"fmt"
	"log"

	"github.com/plivo/plivo-go"
)

func main() {
	client, err := plivo.NewClient("<auth_id>", "<auth_token>", &plivo.ClientOptions{})
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

import "github.com/plivo/plivo-go/plivo/xml"

func main()  {
  println(xml.ResponseElement{
    Contents: []interface{}{
      new(xml.SpeakElement).AddSpeak("Hello, world!", "MAN", "en", 1),
    },
    }.String())
}
```

This generates the following XML:

```xml
<Response>
   <Speak voice="MAN" language="en-us" loop="1">Hello, world!</Speak>
</Response>
```

### Run a PHLO

```go
package main

import (
	"fmt"
	"plivo-go"
)

// Initialize the following params with corresponding values to trigger resources

const authId  = "auth_id"
const authToken = "auth_token"
const phloId = "phlo_id"

// with payload in request

func main() {
	testPhloRunWithParams()
}

func testPhloRunWithParams() {
	phloClient,err := plivo.NewPhloClient(authId, authToken, &plivo.ClientOptions{})
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
		"from": "+14156667778",
		"to": "+14156667778",
	})

	if (err != nil) {
		println(err)
	}
	fmt.Printf("Response: %#v\n", response)
}
```

### More examples
More examples are available [here](https://github.com/plivo/plivo-examples-go). Also refer to the [guides for configuring the Go server to run various scenarios](https://www.plivo.com/docs/sms/quickstart/go-server/) & use it to test out your integration in under 5 minutes.

## Reporting issues
Report any feedback or problems with this version by [opening an issue on Github](https://github.com/plivo/plivo-go/issues).
