package handlertest

import (
	"github.com/krzysztofmadejski/handlertest/internal"
	"net/http"
	"testing"
)

func TestMethod(t *testing.T) {
	expectMethod := func(method string) http.HandlerFunc {
		at := internal.CallerInfo()[1]
		return func(w http.ResponseWriter, r *http.Request) {
			// expect handler to be called with given method
			if r.Method != method {
				t.Errorf("Expected method %s, got %s \nat %v", method, r.Method, at)
			}
		}
	}

	Call(expectMethod("POST")).POST("/jobs").Assert(new(testing.T))
	Call(expectMethod("GET")).GET("/jobs").Assert(new(testing.T))
	Call(expectMethod("OPTIONS")).Method("OPTIONS").Assert(new(testing.T))
}
