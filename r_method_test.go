package handlertest

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMethod(t *testing.T) {
	expectMethod := func(method string) http.HandlerFunc {
		at := assert.CallerInfo()[1]
		return func(w http.ResponseWriter, r *http.Request) {
			// expect handler to be called with given method
			if r.Method != method {
				t.Errorf("Expected method %s, got %s \nat %v", method, r.Method, at)
			}
		}
	}

	Call(expectMethod("POST")).POST("/jobs").Assert(new(testing.T)).Test()
	Call(expectMethod("GET")).GET("/jobs").Assert(new(testing.T)).Test()
	Call(expectMethod("OPTIONS")).Method("OPTIONS").Assert(new(testing.T)).Test()
}
