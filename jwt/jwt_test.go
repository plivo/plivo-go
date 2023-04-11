package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	token := AccessToken{}
	err := token.New("MADADADADADADADADADA", "qwerty", "username", time.Unix(12121212, 0), time.Duration(30)*time.Second, time.Time{}, "username-12345")
	assert.EqualError(t, err, "lifetime out of [180, 86400]")
	_ = token.New("MADADADADADADADADADA", "qwerty", "username", time.Unix(12121212, 0), time.Duration(300)*time.Second, time.Time{}, "username-12345")
	token.AddVoiceGrants(VoiceGrants{true, true})
	assert.Equal(t,
		"eyJhbGciOiJIUzI1NiIsImN0eSI6InBsaXZvO3Y9MSIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJNQURBREFEQURBREFEQURBREFEQSIsInN1YiI6InVzZXJuYW1lIiwiZXhwIjoxMjEyMTUxMiwibmJmIjoxMjEyMTIxMiwianRpIjoidXNlcm5hbWUtMTIzNDUiLCJncmFudHMiOnsidm9pY2UiOnsiaW5jb21pbmdfYWxsb3ciOnRydWUsIm91dGdvaW5nX2FsbG93Ijp0cnVlfX19.IpHETNEfApAE6aalOgb4OazlfFB64PF02oP-ww5TfSM",
		token.ToJwt(),
	)
}
