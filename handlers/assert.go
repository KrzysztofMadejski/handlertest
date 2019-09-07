package handlers

import (
	"net/http"
	"testing"
)

type Assert struct {
	r *Request

	code           int // status code
	headersSet     http.Header
	headersMissing http.Header
}

func NewAssert(r *Request) *Assert {
	return &Assert{
		r:              r,
		headersSet:     make(http.Header),
		headersMissing: make(http.Header),
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

func (a *Assert) Header(key string, value string) *Assert {
	a.headersSet.Set(key, value)
	return a
}

func (a *Assert) HeaderMissing(key string) *Assert {
	a.headersMissing.Set(key, "")
	return a
}

func (a *Assert) ContentType(contentType string) *Assert {
	a.headersSet.Set("Content-Type", contentType)
	return a
}
