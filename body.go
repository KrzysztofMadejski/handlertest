package handlertest

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"strings"
	"testing"
)

func (r *Request) getBodyReader(t *testing.T) io.Reader {
	if r.files != nil {
		// multipart encoding

		var b bytes.Buffer
		var fw io.Writer
		var err error

		w := multipart.NewWriter(&b)
		for field, values := range r.fields {
			for _, v := range values {
				if fw, err = w.CreateFormField(field); err != nil {
					t.Error(err)
				}
				if _, err = fmt.Fprint(fw, v); err != nil {
					t.Error(err)
				}
			}
		}

		for field, files := range r.files {
			for filename, reader := range files {
				if c, isCloser := reader.(io.Closer); isCloser {
					defer func() {
						handle(t, c.Close())
					}()
				}
				if fw, err = w.CreateFormFile(field, filename); err != nil {
					t.Error(err)
				}
				if _, err = io.Copy(fw, reader); err != nil {
					t.Error(err)
				}
			}
		}
		handle(t, w.Close())
		r.ContentType(w.FormDataContentType())

		return &b
	}

	// if json or url-encoded body, just return it
	return strings.NewReader(r.body)
}

func (r *Request) JSON(json string) *Request {
	r.body = json
	return r.ContentType(ContentTypeJSON)
}

func (r *Request) FormURLEncoded(values url.Values) *Request {
	r.body = values.Encode()

	if r.method == "" {
		r.method = "POST"
	}
	return r.ContentType(ContentTypeFormURLEncoded)
}

func (r *Request) FormURLEncodedMap(values map[string]string) *Request {
	return r.FormURLEncoded(ValuesFromMap(values))
}

func (r *Request) FileReaders(fields map[string]map[string]io.Reader) *Request {
	r.files = fields
	return r
}

func (r *Request) Files(fields map[string]map[string]string) *Request {
	flds := make(map[string]map[string]io.Reader)
	for fld, files := range fields {
		readersMap := make(map[string]io.Reader)
		for name, content := range files {
			readersMap[name] = strings.NewReader(content)
		}
		flds[fld] = readersMap
	}

	r.files = flds
	return r
}

func (r *Request) File(field string, fileName string, content string) *Request {
	return r.FileReader(field, fileName, strings.NewReader(content))
}

func (r *Request) FileReader(field string, fileName string, content io.Reader) *Request {
	if r.files == nil {
		r.files = make(map[string]map[string]io.Reader)
	}
	if _, exists := r.files[field]; !exists {
		r.files[field] = make(map[string]io.Reader)
	}
	r.files[field][fileName] = content

	return r
}

func (r *Request) FormMultipart(fields url.Values) *Request {
	if r.files == nil {
		// presence of r.files says it will be multipart form
		r.files = make(map[string]map[string]io.Reader)
	}
	r.fields = fields
	return r
}

func (r *Request) FormMultipartMap(values map[string]string) *Request {
	return r.FormMultipart(ValuesFromMap(values))
}
