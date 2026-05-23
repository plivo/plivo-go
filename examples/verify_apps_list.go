package main

import (
	"fmt"

	plivo "github.com/plivo/plivo-go"
)

func main() {
	client, err := plivo.NewClient("YOUR_AUTH_ID", "YOUR_AUTH_TOKEN", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}

	resp, err := client.VerifyApps.List(plivo.VerifyAppListParams{
		Limit:  20,
		Offset: 0,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Total apps: %d\n", resp.Meta.TotalCount)
	for _, app := range resp.VerifyApps {
		fmt.Printf("App UUID: %s, Name: %s\n", app.AppUUID, app.Name)
	}
}