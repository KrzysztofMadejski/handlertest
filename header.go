package handlertest

func (r *Request) Header(key string, value string) *Request {
	r.headers.Set(key, value)
	return r
}

func (r *Request) ContentType(contentType string) *Request {
	return r.Header("Content-Type", contentType)
}

const ContentTypeFormURLEncoded = "application/x-www-form-urlencoded"
const ContentTypeMultipartFormDataPrefix = "multipart/form-data;"
const ContentTypeJSON = "application/json"
