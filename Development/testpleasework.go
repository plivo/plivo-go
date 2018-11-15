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
	fmt.Println(phloOut)

	fmt.Printf("%# v", pretty.Formatter(phloOut))



}

//https://phlorunnner.plivo.com/v1/phlo/{phlo_id}\
//protected static String BASE_URL = "https://phlorunner.plivo.com/v1/";
//("phlo/{phloId}")
//
//requestUrl = {net/url.URL}
//Scheme = "https"
//Opaque = ""
//*User = {*net/url.Userinfo} nil
//Host = "phlorunner.plivo.com"
//Path = "/v1/phlo"
//RawPath = ""
//ForceQuery = false
//RawQuery = "phlo_id=ebee832f-6467-4e9f-b168-4f296cbc2f09"
////Fragment = ""
//
//
//requestUrl = {net/url.URL}
//Scheme = "https"
//Opaque = ""
//*User = {*net/url.Userinfo} nil
//Host = "api.plivo.com"
//Path = "/v1/Account/MAZMI2NZE5N2EWZDI4MZ/Pricing/"
//RawPath = ""
//ForceQuery = false
//RawQuery = "country_iso=US"
//Fragment = ""