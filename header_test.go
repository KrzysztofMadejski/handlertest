package handlertest

import (
	"github.com/krzysztofmadejski/handlertest/internal"
	"net/http"
	"testing"
)

var expectHeader = func(t *testing.T, header string, expectedValue string) http.HandlerFunc {
	at := internal.CallerInfo()[1]
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
	Call(expectHeader(t, "Allow-Origin", "*")).Header("Allow-Origin", "*").Assert(new(testing.T))
	Call(expectHeader(t, "Allow-Origin", "")).Header("Content-Type", "text/plain").Assert(new(testing.T))
}

func TestContentType(t *testing.T) {
	Call(expectHeader(t, "Content-Type", "text/plain")).ContentType("text/plain").Assert(new(testing.T))
}
