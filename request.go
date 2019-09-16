package handlertest

import (
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	handler http.HandlerFunc

	method  string
	url     string
	headers http.Header
	body    string
	files   map[string]map[string]io.Reader
	fields  url.Values
	custom  func(request *http.Request)

	// TODO context
	// TODO or interface
}

func Call(handler http.HandlerFunc) *Request {
	return &Request{
		handler: handler,
		headers: http.Header{},
	}
}

func (r *Request) Assert() *Assert {
	return NewAssert(r)
}
