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

func TestComputeSignature(t *testing.T) {
	assert.Equal(t,
		"EJEt0ELanhr8hjMPIJnLNLex0dE=",
		ComputeSignature("MAXXXXXXXXXXXXXXXXXX", "http://foo.com/answer/", map[string]string{
			"CallUUID": "97ceeb52-58b6-11e1-86da-77300b68f8bb",
			"Duration": "300",
		}))
}

func TestValidateSignature(t *testing.T) {
	assert.Equal(t, true,
		ValidateSignature("MAXXXXXXXXXXXXXXXXXX", "http://foo.com/answer/", map[string]string{
			"CallUUID": "97ceeb52-58b6-11e1-86da-77300b68f8bb",
			"Duration": "300",
		},
			"EJEt0ELanhr8hjMPIJnLNLex0dE="))
}

func TestComputeSignatureEncoding(t *testing.T) {
	assert.Equal(t, "n3Xfo4u+vRFyl3gsH8B0qDUIK5g=",
		ComputeSignature("MAXXXXXXXXXXXXXXXXXX", "http://foo.com/answer/", map[string]string{
			"a": "1 2",
		}))
}

func TestComputeSignatureFail(t *testing.T) {
	assert.Equal(t, false,
		ValidateSignature("MAXXXXXXXXXXXXXXXXXX", "http://foo.com/answer/", map[string]string{
			"CallUUID": "97ceeb52-58b6-11e1-86da-77300b6b8f8bb",
			"Duration": "300",
		},
			"EJEt0ELanhr8hjMPIJnLNLex0dE="))
}
