package handlertest

func (r *Request) Url(url string) *Request {
	r.url = url
	return r
}
