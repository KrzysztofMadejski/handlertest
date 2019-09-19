package handlertest

func (r *Request) Method(method string) *Request {
	r.method = method
	return r
}

func (r *Request) GET(url string) *Request {
	r.method = "GET"
	return r.Url(url)
}

func (r *Request) POST(url string) *Request {
	r.method = "POST"
	return r.Url(url)
}
