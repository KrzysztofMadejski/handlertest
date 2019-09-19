package handlertest

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"testing"
)

var expectBody = func(t *testing.T, expectedBody string, contentType string) http.HandlerFunc {
	// TODo pull that func in
	at := assert.CallerInfo()[1]
	return func(w http.ResponseWriter, r *http.Request) {
		expectHeader(t, "Content-Type", contentType)(w, r)

		bodyBytes, err := ioutil.ReadAll(r.Body)

		if err != nil {
			t.Errorf("Could not read body because %v", err)

		} else {
			bodyString := string(bodyBytes)
			assert.Equalf(t, expectedBody, bodyString, "Expected body to be different at %s", at)
		}
	}
}

func TestJsonBody(t *testing.T) {
	json := `{"f": 1}`
	Call(expectBody(t, json, ContentTypeJson)).
		Json(json).
		Assert(new(testing.T))
}

var expectForm = func(t *testing.T, expectedValues url.Values, is_multipart bool, method string, numFiles int) http.HandlerFunc {
	// TODo pull that CallerInfo func in
	at := assert.CallerInfo()[1]
	return func(w http.ResponseWriter, r *http.Request) {
		if is_multipart {
			cType := r.Header.Get("Content-Type")
			if !strings.HasPrefix(cType, ContentTypeMultipartFormDataPrefix) {
				t.Errorf("Expected Content-Type to start with '%s', but received '%s'", ContentTypeMultipartFormDataPrefix, cType)

			} else {
				handle(t, r.ParseMultipartForm(1>>24))
			}
		} else { // encoded
			expectHeader(t, "Content-Type", ContentTypeFormUrlEncoded)(w, r)
			handle(t, r.ParseForm())
		}

		assert.Equal(t, expectedValues, r.Form, "Expected request.Form to be populated at %s", at)
		assert.Equal(t, expectedValues, r.PostForm, "Expected request.PostForm to be populated at %s", at)

		if numFiles > 0 {
			if assert.NotNil(t, r.MultipartForm, "Expected form to have files") {
				if assert.NotNil(t, r.MultipartForm.File, "Expected form to have files") {
					for _, fheaders := range r.MultipartForm.File {
						sort.Slice(fheaders, func(i, j int) bool {
							return fheaders[i].Filename < fheaders[j].Filename
						})
						for i, fh := range fheaders {
							assert.Equal(t, fmt.Sprintf("file%d.txt", i+1), fh.Filename)

							f, err := fh.Open()
							if err != nil {
								t.Error(err)
							}
							defer func(f multipart.File) {
								handle(t, f.Close())
							}(f)

							bytes, err := ioutil.ReadAll(f)
							if err != nil {
								t.Error(err)
							}
							assert.Equal(t, fmt.Sprintf("contents%d", i+1), string(bytes), "File content")
						}
					}
				}
			}
		}
	}
}

func TestFormUrlEncoded(t *testing.T) {
	values := url.Values{"field": []string{"val1", "val2"}}
	Call(expectForm(t, values, false, "POST", 0)).
		FormUrlEncoded(values).
		Assert(new(testing.T))
}

// TestFormUrlEncodedSetOtherMethod tests other methods than POST
// POST is set as a default method for sending forms
// but you might set another one
func TestFormUrlEncodedSetOtherMethod(t *testing.T) {
	values := url.Values{"field": []string{"val1", "val2"}}
	Call(expectForm(t, values, false, "PUT", 0)).
		Method("PUT").FormUrlEncoded(values).
		Assert(new(testing.T))

	Call(expectForm(t, values, false, "PUT", 0)).
		FormUrlEncoded(values).Method("PUT").
		Assert(new(testing.T))
}

func TestFormUrlEncodedMap(t *testing.T) {
	values := url.Values{"field": []string{"value"}}
	Call(expectForm(t, values, false, "POST", 0)).
		FormUrlEncodedMap(map[string]string{"field": "value"}).
		Assert(new(testing.T))
}

func TestFormMultipartOneField(t *testing.T) {
	values := url.Values{"field": []string{"val1"}}

	Call(expectForm(t, values, true, "POST", 0)).
		FormMultipart(values).
		Assert(new(testing.T))
}

func TestFormMultipartOnlyFields(t *testing.T) {
	values := url.Values{"field": []string{"val1", "val2"}}

	Call(expectForm(t, values, true, "POST", 0)).
		FormMultipart(values).
		Assert(new(testing.T))
}

func TestFormMultipartMapOnlyFields(t *testing.T) {
	values := url.Values{"field": []string{"value"}}
	Call(expectForm(t, values, true, "POST", 0)).
		FormMultipartMap(map[string]string{"field": "value"}).
		Assert(new(testing.T))
}

func TestFormMultipartOneFile(t *testing.T) {
	Call(expectForm(t, url.Values{}, true, "POST", 1)).
		Files(map[string]map[string]string{"files[]": {"file1.txt": "contents1"}}).
		Assert(new(testing.T))
}

func TestFormMultipartMultipleFiles(t *testing.T) {
	Call(expectForm(t, url.Values{}, true, "POST", 2)).
		Files(map[string]map[string]string{"files[]": {"file1.txt": "contents1", "file2.txt": "contents2"}}).
		Assert(new(testing.T))
}

func TestFormMultipartMultipleFileReaders(t *testing.T) {
	Call(expectForm(t, url.Values{}, true, "POST", 2)).
		FileReaders(map[string]map[string]io.Reader{"files[]": {"file1.txt": strings.NewReader("contents1"), "file2.txt": strings.NewReader("contents2")}}).
		Assert(new(testing.T))
}

func TestFormMultipartAddFileReader(t *testing.T) {
	Call(expectForm(t, url.Values{}, true, "POST", 1)).
		FileReader("files[]", "file1.txt", strings.NewReader("contents1")).
		Assert(new(testing.T))
}

func TestFormMultipartAddFile(t *testing.T) {
	Call(expectForm(t, url.Values{}, true, "POST", 1)).
		File("files[]", "file1.txt", "contents1").
		Assert(new(testing.T))
}

func TestFormMultipartAddFileMultiple(t *testing.T) {
	Call(expectForm(t, url.Values{}, true, "POST", 2)).
		File("files[]", "file1.txt", "contents1").
		File("files[]", "file2.txt", "contents2").
		Assert(new(testing.T))
}

// TODO test joining url values and form values
