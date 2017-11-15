package plivo

import "testing"
import (
	"github.com/stretchr/testify/assert"
)

func TestNumbers(t *testing.T) {
	num1 := "+911231231230"
	num2 := "+913213213210"
	joined := "+911231231230<+913213213210"
	assert.Equal(t, joined, Numbers(num1, num2))
}

func TestHeaders(t *testing.T) {
	headers := map[string]string{
		"X-PH-Test1": "value1",
		"X-PH-Test2": "value2",
	}
	encoded := Headers(headers)
	// Go iterates over maps in random order
	assert.Contains(t, []string{"X-PH-Test1=value1,X-PH-Test2=value2", "X-PH-Test2=value2,X-PH-Test1=value1"}, encoded)
}

func TestComputeSignaturePass(t *testing.T) {
	assert.Equal(t, "ehV3IKhLysWBxC1sy8INm0qGoQYdYsHwuoKjsX7FsXc=",ComputeSignature(
		"my_auth_token",
		"https://answer.url",
		"12345"
		)
	)
}

func TestComputeSignatureFail(t *testing.T) {
	assert.Equal(t, "ehV3IKhLysWBxC1sy8INm0qGoQYdYsHwuoKjsX7FsXc=",ComputeSignature(
		"my_auth_tokens",
		"https://answer.url",
		"12345"
		)
	)
}

func TestValidateSignaturePass(t *testing.T) {
	assert.Equal(t, true,ValidateSignatureV2(
				"https://answer.url",
				"12345",
				"ehV3IKhLysWBxC1sy8INm0qGoQYdYsHwuoKjsX7FsXc=",
				"my_auth_token"
			)
	)
}
func TestValidateSignatureV2Fail(t *testing.T) {
	assert.Equal(t, false,ValidateSignatureV2(
				"https://answer.url",
				"12345",
				"ehV3IKhLysWBxC1sy8INm0qGoQYdYsHwuoKjsX7FsXc=",
				"my_auth_tokens"
			)
	)
}
