package main

import (
	"log"
	"time"

	"github.com/plivo/plivo-go/v7/jwt"
)

func mainJWT() {
	token1 := jwt.AccessToken{}
	err := token1.New("{authId}", "{authToken}", "{endpointUsername}", time.Now(), time.Duration(300)*time.Second, time.Time{}, "{uid}")
	if err != nil {
		log.Fatalf("abort: %+v\n", err)
	}
	token1.AddVoiceGrants(jwt.VoiceGrants{IncomingAllow: false, OutgoingAllow: true})
	log.Println(token1.ToJwt())

	token2 := jwt.AccessToken{}
	err = token2.New("{authId}", "{authToken}", "{endpointUsername}", time.Now(), 0, time.Now().Add(300*time.Second), "{uid}")
	if err != nil {
		log.Fatalf("abort: %+v\n", err)
	}
	token2.AddVoiceGrants(jwt.VoiceGrants{IncomingAllow: false, OutgoingAllow: true})
	log.Println(token2.ToJwt())
}
