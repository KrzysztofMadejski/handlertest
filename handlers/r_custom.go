package handlers

import "net/http"

func (r *Request) Custom(customize func(request *http.Request)) *Request {
	r.custom = customize
	return r
}
