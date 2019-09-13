package handlers

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

	// TODO context
	// TODO or interface
}

func NewRequest(handler http.HandlerFunc) *Request {
	return &Request{
		handler: handler,
		headers: http.Header{}, // TODO will make initialize that?
	}
}

func (r *Request) Assert() *Assert {
	return NewAssert(r)
}
