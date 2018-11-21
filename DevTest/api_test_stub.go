package main

import (
	"fmt"
	"github.com/kr/pretty"
	"plivo-go"
)

func main()  {
	client, err := plivo.NewPhloClient("MAZMI2NZE5N2EWZDI4MZ", "NzUyYTVhMTY5MDczZmRjNDk1NmI5YTJmNTgwMDI4", &plivo.ClientOptions{})

	phloOut , err := client.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	if err != nil {
		panic(err)
	}
//https://phlorunner.plivo.com/v1/account/MAZMI2NZE5N2EWZDI4MZ/phlo/2e802f94-cc03-4e35-b4c8-5a0328544135

	type PhloRunPayLoad struct {
		To string `to:"to" url:"to"`
	}

	//prpl := PhloRunPayLoad{"+919503364736"}
	type runnerParams map[string]interface{}
	phloRun, err := phloOut.Run(runnerParams{"to":"+919503364736" })

	fmt.Printf("%# v", pretty.Formatter(phloRun))



	mout,err := phloOut.MultiPartyCall("2e802f94-cc03-4e35-b4c8-5a0328544135","multi_party_call","b18848dd-e308-4758-bc5c-fb9c274d5fb7")
	fmt.Printf("MultiPartyCall output %# v", pretty.Formatter(mout))
	if err != nil {
			panic(err)
		}



	mpcap := plivo.MultiPartyCallActionPayload{"call","+919441456465","agent","+919503364736"}


	nodeOut,err := mout.Call(mpcap)
	fmt.Printf(" nodeOut %# v", pretty.Formatter(nodeOut))



}
