package handlertest

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var expectHeader = func(t *testing.T, header string, expectedValue string) http.HandlerFunc {
	at := assert.CallerInfo()[1]
	return func(w http.ResponseWriter, r *http.Request) {
		value := r.Header.Get(header)
		if expectedValue != "" && value == "" {
			t.Errorf("Expected Header %s set to  %s, but it is empty \nat %v", header, expectedValue, at)

		} else if expectedValue == "" && value != "" {
			t.Errorf("Expected Header %s to be empty, but got %s \nat %v", header, value, at)

		} else if expectedValue != value {
			t.Errorf("Expected Header %s set to %s, got %s \nat %v", header, expectedValue, value, at)
		}
	}
}

func TestHeader(t *testing.T) {
	Call(expectHeader(t, "Allow-Origin", "*")).Header("Allow-Origin", "*").Assert(new(testing.T)).Test()
	Call(expectHeader(t, "Allow-Origin", "")).Header("Content-Type", "text/plain").Assert(new(testing.T)).Test()
}

func TestContentType(t *testing.T) {
	Call(expectHeader(t, "Content-Type", "text/plain")).ContentType("text/plain").Assert(new(testing.T)).Test()
}
