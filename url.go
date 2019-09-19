package handlertest

func (r *Request) URL(url string) *Request {
	r.url = url
	return r
}
