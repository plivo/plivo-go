package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	token := AccessToken{}
	err := token.New("MADADADADADADADADADA", "qwerty", "username", time.Unix(12121212, 0), time.Duration(30) * time.Second, time.Time{}, "username-12345")
	assert.EqualError(t, err, "lifetime out of [180, 84600]")
	_ = token.New("MADADADADADADADADADA", "qwerty", "username", time.Unix(12121212, 0), time.Duration(300) * time.Second, time.Time{}, "username-12345")
	token.AddVoiceGrants(VoiceGrants{true, true})
	assert.Equal(t,
		"eyJhbGciOiJIUzI1NiIsImN0eSI6InBsaXZvO3Y9MSIsInR5cCI6IkpXVCJ9.eyJleHAiOjEyMTIxNTEyLCJqdGkiOiJ1c2VybmFtZS0xMjM0NSIsImlzcyI6Ik1BREFEQURBREFEQURBREFEQURBIiwibmJmIjoxMjEyMTIxMiwic3ViIjoidXNlcm5hbWUiLCJncmFudHMiOnsidm9pY2UiOnsiaW5jb21pbmdfYWxsb3ciOnRydWUsIm91dGdvaW5nX2FsbG93Ijp0cnVlfX19.51cD6eyvu_wvQ_3ag2i_8_T4cvbkU__rnkyH2MBA5Nk",
		token.ToJwt(),
	)
}
