package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type Assert struct {
	r *Request

	code int // status code
}

func (a *Assert) Test(t *testing.T) {
	// TODO move request creation to request.go?

	// set method & url
	req, err := http.NewRequest(a.r.method, a.r.url, a.r.getBody())
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

	//req.URL.RawQuery = values.Encode()


	recorder := httptest.NewRecorder()

	a.r.handler.ServeHTTP(recorder, req)

	// tests below
	if a.code > 0 && recorder.Code != a.code {
		t.Errorf("Expected statusCode %d, got %d", a.code, recorder.Code)
	}
}

func (a *Assert) TestRun() func(*testing.T) {
	return func(t *testing.T) {
		a.Test(t)
	}
}

func (a *Assert) Status(statusCode int) *Assert {
	a.code = statusCode

	return a
}