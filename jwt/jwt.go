package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AccessToken struct {
	AuthId    string
	Key       string
	Username  string
	ValidFrom time.Time
	Lifetime  time.Duration
	Uid       string
	Grants    Grants `json:"grants"`
}

type Grants struct {
	Voice VoiceGrants `json:"voice"`
}

type VoiceGrants struct {
	IncomingAllow bool `json:"incoming_allow"`
	OutgoingAllow bool `json:"outgoing_allow"`
}

type JwtClaims struct {
	jwt.StandardClaims
	Grants Grants `json:"grants"`
}

func (acctkn *AccessToken) New(authId string, authToken string, username string, validFrom time.Time, lifetime time.Duration, validTill time.Time, uid string) error {
	if len(authId) == 0 {
		authId = os.Getenv("PLIVO_AUTH_ID")
	}
	if len(authId) == 0 {
		return errors.New("authId not found")
	}
	acctkn.AuthId = authId
	if len(authToken) == 0 {
		authToken = os.Getenv("PLIVO_AUTH_TOKEN")
	}
	if len(authToken) == 0 {
		return errors.New("authToken not found")
	}
	acctkn.Key = authToken
	if len(username) == 0 {
		return errors.New("username not found")
	}
	acctkn.Username = username

	if lifetime != 0 && (lifetime < 180*time.Second || lifetime > 86400*time.Second) {
		return errors.New("lifetime out of [180, 86400]")
	}

	if validFrom.IsZero() {
		acctkn.Lifetime = lifetime
		if validTill.IsZero() {
			acctkn.ValidFrom = time.Now()
		} else {
			acctkn.ValidFrom = validTill.Add(-acctkn.Lifetime)
		}
	} else if validTill.IsZero() {
		acctkn.Lifetime = lifetime
		acctkn.ValidFrom = validFrom

	} else {
		if lifetime != 0 {
			return errors.New("use any 2 of validFrom, lifetime and validTill")
		}
		acctkn.ValidFrom = validFrom
		acctkn.Lifetime = validTill.Sub(validFrom)
	}

	acctkn.Uid = uid
	if len(uid) == 0 {
		acctkn.Uid = fmt.Sprintf("%s-%d", acctkn.Username, time.Now().UnixNano())
	}

	return nil
}

func (acctkn *AccessToken) AddVoiceGrants(grants VoiceGrants) {
	acctkn.Grants = Grants{grants}
}

func (acctkn *AccessToken) ToJwt() string {
	// Create the Claims
	claims := JwtClaims{
		jwt.StandardClaims{
			Id:        acctkn.Uid,
			Issuer:    acctkn.AuthId,
			Subject:   acctkn.Username,
			NotBefore: acctkn.ValidFrom.Unix(),
			ExpiresAt: acctkn.ValidFrom.Add(acctkn.Lifetime).Unix(),
		},
		acctkn.Grants,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["cty"] = "plivo;v=1"
	ss, _ := token.SignedString([]byte(acctkn.Key))
	return ss
}
