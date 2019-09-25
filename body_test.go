package handlertest

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/krzysztofmadejski/handlertest/internal"
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
	at := internal.CallerInfo()[1]
	return func(w http.ResponseWriter, r *http.Request) {
		expectHeader(t, "Content-Type", contentType)(w, r)

		bodyBytes, err := ioutil.ReadAll(r.Body)

		if err != nil {
			t.Errorf("Could not read body because %v", err)

		} else {
			bodyString := string(bodyBytes)
			if expectedBody != bodyString {
				t.Errorf("Expected body to be different at %s. Expected '%s', got '%s'", at, expectedBody, bodyString)
			}
		}
	}
}

func TestJsonBody(t *testing.T) {
	json := `{"f": 1}`
	Call(expectBody(t, json, ContentTypeJSON)).
		JSON(json).
		Assert(new(testing.T))
}

var expectForm = func(t *testing.T, expectedValues url.Values, is_multipart bool, method string, numFiles int) http.HandlerFunc {
	at := internal.CallerInfo()[1]
	return func(w http.ResponseWriter, r *http.Request) {
		if is_multipart {
			cType := r.Header.Get("Content-Type")
			if !strings.HasPrefix(cType, ContentTypeMultipartFormDataPrefix) {
				t.Errorf("Expected Content-Type to start with '%s', but received '%s'", ContentTypeMultipartFormDataPrefix, cType)

			} else {
				handle(t, r.ParseMultipartForm(1>>24))
			}
		} else { // encoded
			expectHeader(t, "Content-Type", ContentTypeFormURLEncoded)(w, r)
			handle(t, r.ParseForm())
		}

		if !cmp.Equal(expectedValues, r.Form) {
			t.Errorf("Expected request.Form to be %+v. but got %+v at %s", expectedValues, r.Form, at)
		}
		if !cmp.Equal(expectedValues, r.PostForm) {
			t.Errorf("Expected request.PostForm to be %+v. but got %+v at %s", expectedValues, r.Form, at)
		}

		if numFiles == 0 {
			return
		}
		if r.MultipartForm == nil || r.MultipartForm.File == nil {
			t.Errorf("Expected MultipartForm to be set and have files")
			return
		}
		for _, fheaders := range r.MultipartForm.File {
			sort.Slice(fheaders, func(i, j int) bool {
				return fheaders[i].Filename < fheaders[j].Filename
			})
			for i, fh := range fheaders {
				fileName := fmt.Sprintf("file%d.txt", i+1)
				if fileName != fh.Filename {
					t.Errorf("Expected file %s at position %d", fileName, i)
				}

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
				actualContents := string(bytes)
				expectedContents := fmt.Sprintf("contents%d", i+1)
				if expectedContents != actualContents {
					t.Errorf("Expected content '%s', but got '%s' at %d", expectedContents, actualContents, i)
				}
			}
		}
	}
}

func TestFormUrlEncoded(t *testing.T) {
	values := url.Values{"field": []string{"val1", "val2"}}
	Call(expectForm(t, values, false, "POST", 0)).
		FormURLEncoded(values).
		Assert(new(testing.T))
}

// TestFormUrlEncodedSetOtherMethod tests other methods than POST
// POST is set as a default method for sending forms
// but you might set another one
func TestFormUrlEncodedSetOtherMethod(t *testing.T) {
	values := url.Values{"field": []string{"val1", "val2"}}
	Call(expectForm(t, values, false, "PUT", 0)).
		Method("PUT").FormURLEncoded(values).
		Assert(new(testing.T))

	Call(expectForm(t, values, false, "PUT", 0)).
		FormURLEncoded(values).Method("PUT").
		Assert(new(testing.T))
}

func TestFormUrlEncodedMap(t *testing.T) {
	values := url.Values{"field": []string{"value"}}
	Call(expectForm(t, values, false, "POST", 0)).
		FormURLEncodedMap(map[string]string{"field": "value"}).
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
