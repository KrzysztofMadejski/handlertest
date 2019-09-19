package assert

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func (a *Assert) Test() {
	response := a.r
	t := a.t

	// test status code
	if a.code > 0 && response.StatusCode != a.code {
		t.Errorf("Expected statusCode %d, got %d", a.code, response.StatusCode)
	}

	// test headers
	for key, expectedValues := range a.headersSet {
		values := response.Header[key]
		if values == nil || len(values) == 0 {
			t.Errorf("Expected header %s to be set, it is not", key)

		} else {
			assert.Equalf(t, expectedValues, values, "Expected header %s to be set to different value", key)
		}
	}

	for key := range a.headersMissing {
		values := response.Header[key]
		assert.Empty(t, values, "Expected header %s to be empty", key)
	}

	// Test body
	if a.body != nil {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Errorf("Could not read response body: %v", err)
		}

		a.body(t, body)
	}

	// Run any custom assertions
	if a.custom != nil {
		a.custom(t, response)
	}
}

//noinspection GoUnusedExportedFunction
func TestThisFileIsNotATest(t *testing.T) {
	t.Fatalf("This method should not be executed during tests")
}
