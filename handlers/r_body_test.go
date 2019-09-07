package handlers

import (
	"testing"
)

//func TestJsonBody(t *testing.T) {
//	expectMethod := func(method string) http.HandlerFunc {
//		at := assert.CallerInfo()[1]
//		return func(w http.ResponseWriter, r *http.Request) {
//			// expect handler to be called with given method
//			if r.Method != method {
//				t.Errorf("Expected method %s, got %s \nat %v", method, r.Method, at)
//			}
//		}
//	}
//
//	NewRequest(expectMethod("POST")).POST("/jobs").Assert().Test(new(testing.T))
//}
// application/x-www-form-urlencoded
// multipart/form-data

func TestFormUrlEncoded(t *testing.T) {
	t.Skip("Not implemented")
}

func TestFormMultipart(t *testing.T) {
	t.Skip("Not implemented")
}

func TestFormMultipartFiles(t *testing.T) {
	t.Skip("Not implemented")
}

// TODO test multiple body options conflicts