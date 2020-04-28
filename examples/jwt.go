package main

import (
	"log"
	"time"

	"../jwt"
)

func main() {
	token := jwt.AccessToken{}
	err := token.New("{authId}", "{authToken}", "{endpointUsername}", time.Now(), time.Duration(300)*time.Second, time.Time{}, "username-12345")
	token.AddVoiceGrants(jwt.VoiceGrants{false, true})
	if err != nil {
		log.Fatalf("abort: %+v\n", err)
	}
	log.Println(token.ToJwt())

}
