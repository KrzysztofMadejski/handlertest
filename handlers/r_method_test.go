package handlers

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

	NewRequest(expectMethod("POST")).POST("/jobs").Assert().Test(new(testing.T))
	NewRequest(expectMethod("GET")).GET("/jobs").Assert().Test(new(testing.T))
	NewRequest(expectMethod("OPTIONS")).Method("OPTIONS").Assert().Test(new(testing.T))
}
