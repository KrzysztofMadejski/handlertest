package assert_test

import (
	"encoding/json"
	"github.com/krzysztofmadejski/handlertest"
	"net/http"
	"testing"
)

// TODO contentType unused
func setBody(content string, contentType string) func(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(content))
		// TODO test for error
	}
	return handler
}

func TestExpectsBodyFunction(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(setBody(`[{"id": 201809}]`, handlertest.ContentTypeJSON)).Assert(mockT).
		Body(func(t *testing.T, body []byte) {
			// don'T raise error on mockT
		})
	if mockT.Failed() {
		t.Errorf("Expected assertion to pass")
	}
}

type Obj struct {
	Id int `json:"id"`
}

func TestExpectsBodyFunctionFails(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(setBody(`[{"id": 201809}]`, handlertest.ContentTypeJSON)).Assert(mockT).
		Body(func(t *testing.T, body []byte) {
			var o Obj
			if err := json.Unmarshal(body, &o); err != nil {
				t.Errorf("Could not unmarshall body")
				return
			}
			if o.Id > 201807 {
				mockT.Errorf("Expected Id to be something it wasn'T")
			}
		})
	if !mockT.Failed() {
		t.Errorf("Expected assertion to fail")
	}
}

// TODO charset: utf-8 in Content-Type
func TestExpectsJsonBody(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(setBody(`[]`, handlertest.ContentTypeJSON)).Assert(mockT).
		JsonBody(`[]`)
	if mockT.Failed() {
		t.Errorf("Expected assertion to pass")
	}
}

func TestExpectsJsonBodyFails(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(setBody(`[]`, handlertest.ContentTypeJSON)).Assert(mockT).
		JsonBody(`[{"id": 1}]`)
	if !mockT.Failed() {
		t.Errorf("Assertion should fail when body is different")
	}
}

// TODO test json indenting
// TODO test error message returns a nice diff during extended run

func TestExpectJsonType(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(setBody(`[{"id": 1}]`, handlertest.ContentTypeJSON)).Assert(mockT).
		JsonUnmarshallsTo([]Obj{})
	if mockT.Failed() {
		t.Errorf("Expected assertion to pass")
	}
}

func TestExpectJsonTypeFails(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(setBody(`{"id": 1}`, handlertest.ContentTypeJSON)).Assert(mockT).
		JsonUnmarshallsTo([]Obj{})
	if !mockT.Failed() {
		t.Errorf("Expected assertion to fail")
	}
}

func TestExpectJsonMatches(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(setBody(`[{"id": 1}]`, handlertest.ContentTypeJSON)).Assert(mockT).
		JsonMatches(func(t *testing.T, list []Obj) {
			if len(list) != 1 {
				t.Errorf("Expected length 0")
			}
			if len(list) < 1 || list[0].Id != 1 {
				t.Errorf("Expected list[0].id=1")
			}
		})

	if mockT.Failed() {
		t.Errorf("Expected assertion to pass")
	}
}

func TestExpectJsonMatchesCantUnmarshall(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(setBody(`[{"id": 1}]`, handlertest.ContentTypeJSON)).Assert(mockT).
		JsonMatches(func(t *testing.T, obj Obj) {})
	if !mockT.Failed() {
		t.Errorf("Expected assertion to fail")
	}
}

func TestExpectJsonMatchesFails(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(setBody(`[{"id": 1}]`, handlertest.ContentTypeJSON)).Assert(mockT).
		// TODO allow to use pointers also JsonMatches(func(T *testing.T, list *[]Obj) {
		JsonMatches(func(t *testing.T, list []Obj) {
			t.Errorf("Fail because something didn'T meet your expectations")
		})
	if !mockT.Failed() {
		t.Errorf("Expected assertion to fail")
	}
}

func TestExpectJsonMatchesWrongFunc(t *testing.T) {
	type test struct {
		name     string
		function interface{}
	}
	for _, tt := range []test{
		{"Empty func", func() {}},

		{"T should be the first arg", func(Obj, Obj) {}},
		{"T should be a pointer", func(testing.T, Obj) {}}, // let vet complain about mutex here
		{"Too many params", func(*testing.T, Obj, Obj) {}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			mockT := new(testing.T)
			handlertest.Call(setBody(`[{"id": 1}]`, handlertest.ContentTypeJSON)).Assert(mockT).
				JsonMatches(tt.function)
			if !mockT.Failed() {
				t.Errorf("Expected assertion to fail")
			}
		})
	}
}

// TODO ConformsToFails on unexported declaration

// TODO test json requests if ContentType not set properly
