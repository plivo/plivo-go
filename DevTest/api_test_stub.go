package main

import (
	"fmt"
	"github.com/kr/pretty"
	"plivo-go"
)

func main()  {
	client, _ := plivo.NewPhloClient("MAZMI2NZE5N2EWZDI4MZ", "NzUyYTVhMTY5MDczZmRjNDk1NmI5YTJmNTgwMDI4", &plivo.ClientOptions{})
	phloOut , _ := client.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	type params map[string]interface{}
	//phloRun1, _ := phloOut.Run(params{"from":"919441456465", "to":"919503364736"})
	//phloRun, _ := phloOut.Run(nil)
	//fmt.Printf("Phlo Run OutPut %# v", pretty.Formatter(phloRun))
	//fmt.Printf("Phlo Error OutPut %# v", pretty.Formatter(phloRun1))


	//payload := plivo.MultiPartyCallActionPayload{"call","919704583677","agent","919503364736"}
	multipartycall,_ := phloOut.MultiPartyCall("2e802f94-cc03-4e35-b4c8-5a0328544135","multi_party_call","b18848dd-e308-4758-bc5c-fb9c274d5fb7")
	//nodeOut,_ := multipartycall.Call(payload)


	//fmt.Printf("MultiPartyCall output %# v", pretty.Formatter(multipartycall))
	//fmt.Printf(" NodeResponse output %# v", pretty.Formatter(nodeOut))

	memberPayload := plivo.MultiPartyCallMemberActionPayload{"hold"}
	memberUpdateCall,_ := multipartycall.Member("919704583677").Hold(memberPayload)

	fmt.Printf(" NodeResponse output %# v", pretty.Formatter(memberUpdateCall))

}
