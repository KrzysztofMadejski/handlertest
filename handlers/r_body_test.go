package handlers

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
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
	NewRequest(expectBody(t, json, ContentTypeJson)).
		Json(json).
		Assert().Test(new(testing.T))
}

var expectForm = func(t *testing.T, expectedValues url.Values, contentType string, method string) http.HandlerFunc {
	// TODo pull that func in
	at := assert.CallerInfo()[1]
	return func(w http.ResponseWriter, r *http.Request) {
		expectHeader(t, "Content-Type", contentType)(w, r)

		r.ParseForm()

		if r.Header.Get("Content-Type") == ContentTypeMultipartFormData {
			r.ParseMultipartForm(1 >> 24)
		}

		assert.Equal(t, expectedValues, r.Form, "Expected request.Form to be populated at %s", at)
		assert.Equal(t, expectedValues, r.PostForm, "Expected request.PostForm to be populated at %s", at)
	}
}

func TestFormUrlEncoded(t *testing.T) {
	values := url.Values{"field": []string{"val1", "val2"}}
	NewRequest(expectForm(t, values, ContentTypeFormUrlEncoded, "POST")).
		FormUrlEncoded(values).
		Assert().Test(new(testing.T))
}

// TestFormUrlEncodedSetOtherMethod tests other methods than POST
// POST is set as a default method for sending forms
// but you might set another one
func TestFormUrlEncodedSetOtherMethod(t *testing.T) {
	values := url.Values{"field": []string{"val1", "val2"}}
	NewRequest(expectForm(t, values, ContentTypeFormUrlEncoded, "PUT")).
		Method("PUT").FormUrlEncoded(values).
		Assert().Test(new(testing.T))

	NewRequest(expectForm(t, values, ContentTypeFormUrlEncoded, "PUT")).
		FormUrlEncoded(values).Method("PUT").
		Assert().Test(new(testing.T))
}

func TestFormUrlEncodedMap(t *testing.T) {
	values := url.Values{"field": []string{"value"}}
	NewRequest(expectForm(t, values, ContentTypeFormUrlEncoded, "POST")).
		FormUrlEncodedMap(map[string]string{"field": "value"}).
		Assert().Test(new(testing.T))
}

func TestFormMultipart(t *testing.T) {
	t.SkipNow() // TODO Multipart should contain boundary string

	values := url.Values{"field": []string{"val1", "val2"}}

	NewRequest(expectForm(t, values, ContentTypeMultipartFormData, "POST")).
		FormMultipart(values).
		Assert().Test(new(testing.T))
}

func TestFormMultipartMap(t *testing.T) {
	t.SkipNow() // TODO Multipart should contain boundary string

	values := url.Values{"field": []string{"value"}}
	NewRequest(expectForm(t, values, ContentTypeMultipartFormData, "POST")).
		FormMultipartMap(map[string]string{"field": "value"}).
		Assert().Test(new(testing.T))
}

// TODO test joining url values and form values

//func TestFormMultipartFiles(t *testing.T) {
//	// TODO
//	// NewRequest(expectBody(t)).FormMultipartMap(map[string]string{"field": "how to pass file?"}).Assert().Test(new(testing.T))
//}
