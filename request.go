package handlertest

import (
	"github.com/krzysztofmadejski/handlertest/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type Request struct {
	t       *testing.T
	handler http.HandlerFunc

	method  string
	url     string
	headers http.Header
	body    string
	files   map[string]map[string]io.Reader
	fields  url.Values
	custom  func(request *http.Request)

	// TODO context
}

func Call(handler http.HandlerFunc) *Request {
	return &Request{
		//t: t, TODO
		handler: handler,
		headers: http.Header{},
	}
}

func (r *Request) Assert(t *testing.T) *assert.Assert {
	return &assert.Assert{R: r.createResponse(t), T: t}
}

func (r *Request) createResponse(t *testing.T) *http.Response {
	// set method & url
	req, err := http.NewRequest(r.method, r.url, r.getBodyReader(t))
	if err != nil {
		t.Fatal(err.Error())
	}

	// set headers
	req.Header = r.headers

	// TODO Populate the request's context with our test data.
	//ctx := req.Context()
	//ctx = context.WithValue(ctx, "app.auth.token", "abc123")
	//ctx = context.WithValue(ctx, "app.user",
	//	&YourUser{ID: "qejqjq", Email: "user@example.com"})
	//
	//// Add our context to the request: note that WithContext returns a copy of
	//// the request, which we must assign.
	//req = req.WithContext(ctx)

	if r.custom != nil {
		r.custom(req)
	}

	// ============================
	// =========== TESTS ==========

	recorder := httptest.NewRecorder()

	r.handler.ServeHTTP(recorder, req)

	return recorder.Result()
}
