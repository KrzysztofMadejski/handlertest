package handlers

func (r *Request) Header(key string, value string) *Request {
	r.headers.Set(key, value)
	return r
}

func (r *Request) ContentType(contentType string) *Request {
	return r.Header("Content-Type", contentType)
}

const ContentTypeFormUrlEncoded = "application/x-www-form-urlencoded"
const ContentTypeMultipartFormData = "multipart/form-data"
const ContentTypeJson = "application/json"
