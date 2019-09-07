package handlers

import (
	"github.com/krzysztofmadejski/testify-handlers/utils"
	"io"
	"net/url"
	"strings"
)

func (r *Request) getBodyReader() io.Reader {
	return strings.NewReader(r.body)
}

func (r *Request) Json(json string) *Request {
	r.body = json
	return r.ContentType(ContentTypeJson)
}

func (r *Request) FormUrlEncoded(values url.Values) *Request {
	r.body = values.Encode()
	if r.method == "" {
		r.method = "POST"
	}
	return r.ContentType(ContentTypeFormUrlEncoded)
}

func (r *Request) FormUrlEncodedMap(values map[string]string) *Request {
	return r.FormUrlEncoded(utils.ValuesFromMap(values))
}

func (r *Request) FormMultipart(values url.Values) *Request {
	if r.method == "" {
		r.method = "POST"
	}
	return r.ContentType(ContentTypeMultipartFormData)
}

func (r *Request) FormMultipartMap(values map[string]string) *Request {
	return r.FormMultipart(utils.ValuesFromMap(values))
}
