package main

import (
	"fmt"
	"github.com/kr/pretty"
	"plivo-go"
)

func main()  {
	client, err := plivo.NewClient("MAZMI2NZE5N2EWZDI4MZ", "NzUyYTVhMTY5MDczZmRjNDk1NmI5YTJmNTgwMDI4", &plivo.ClientOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//price , err:= client.Pricing.Get("US")
	//
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%# v", pretty.Formatter(price))



	phloOut , err := client.Phlo.Get("ebee832f-6467-4e9f-b168-4f296cbc2f09")
	if err != nil {
		panic(err)
	}
	//fmt.Println(phloOut)
	//
	//fmt.Printf("%# v", pretty.Formatter(phloOut))

	nodeOut,err := phloOut.GetNode("ebee832f-6467-4e9f-b168-4f296cbc2f09","multi_party_call","8293c88b-afe2-4a05-988b-3f85668c4a67")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%# v", pretty.Formatter(nodeOut))

}

