package handlers

import (
	"net/http"
)

type Request struct {
	handler http.HandlerFunc

	// TODO how do we keep set options?
	// 1) in struct: a.r.method  , but then it is hard to make it nill if needed..
	// 2) in general map[string]interface{}, but then it is harder to access a.r.map["method"] or a.r["method"]
	method  string
	url     string
	headers http.Header
	body    string

	// TODO context
	// TODO or interface
}

func NewRequest(handler http.HandlerFunc) *Request {
	return &Request{
		handler: handler,
		headers: http.Header{}, // make(map[string]string),
	}
}

func (r *Request) Assert() *Assert {
	return NewAssert(r)
}
