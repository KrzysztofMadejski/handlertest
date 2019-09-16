package handlers

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func (a *Assert) Test(t *testing.T) {
	// TODO move request creation to request.go?

	// set method & url
	req, err := http.NewRequest(a.r.method, a.r.url, a.r.getBodyReader(t))
	if err != nil {
		t.Fatal(err.Error())
	}

	// set headers
	req.Header = a.r.headers

	// TODO Populate the request's context with our test data.
	//ctx := req.Context()
	//ctx = context.WithValue(ctx, "app.auth.token", "abc123")
	//ctx = context.WithValue(ctx, "app.user",
	//	&YourUser{ID: "qejqjq", Email: "user@example.com"})
	//
	//// Add our context to the request: note that WithContext returns a copy of
	//// the request, which we must assign.
	//req = req.WithContext(ctx)

	if a.r.custom != nil {
		a.r.custom(req)
	}

	// ============================
	// =========== TESTS ==========

	recorder := httptest.NewRecorder()

	a.r.handler.ServeHTTP(recorder, req)

	response := recorder.Result()

	// test status code
	if a.code > 0 && recorder.Code != a.code {
		t.Errorf("Expected statusCode %d, got %d", a.code, recorder.Code)
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
