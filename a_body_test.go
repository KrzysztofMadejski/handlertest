package handlertest

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
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
	Call(setBody(`[{"id": 201809}]`, ContentTypeJson)).Assert(mockT).
		Body(func(t *testing.T, body []byte) {
			// don't raise error on mockT
		}).Test()
	assert.False(t, mockT.Failed())
}

type Obj struct {
	Id int `json:"id"`
}

func TestExpectsBodyFunctionFails(t *testing.T) {
	mockT := new(testing.T)
	Call(setBody(`[{"id": 201809}]`, ContentTypeJson)).Assert(mockT).
		Body(func(t *testing.T, body []byte) {
			var o Obj
			if err := json.Unmarshal(body, o); err != nil {
				t.Errorf("Could not unmarshall body")
				return
			}
			if o.Id > 201807 {
				mockT.Errorf("Expected Id to be something it wasn't")
			}
		}).Test()
	assert.True(t, mockT.Failed())
}

// TODO charset: utf-8 in Content-Type
func TestExpectsJsonBody(t *testing.T) {
	mockT := new(testing.T)
	Call(setBody(`[]`, ContentTypeJson)).Assert(mockT).
		JsonBody(`[]`).
		Test()
	assert.False(t, mockT.Failed())
}

func TestExpectsJsonBodyFails(t *testing.T) {
	mockT := new(testing.T)
	Call(setBody(`[]`, ContentTypeJson)).Assert(mockT).
		JsonBody(`[{"id": 1}]`).
		Test()
	assert.True(t, mockT.Failed(), "Assertion should fail when body is different")
}

// TODO test json indenting
// TODO test error message returns a nice diff during extended run

func TestExpectJsonType(t *testing.T) {
	mockT := new(testing.T)
	Call(setBody(`[{"id": 1}]`, ContentTypeJson)).Assert(mockT).
		JsonUnmarshallsTo([]Obj{}).
		Test()
	assert.False(t, mockT.Failed())
}

func TestExpectJsonTypeFails(t *testing.T) {
	mockT := new(testing.T)
	Call(setBody(`{"id": 1}`, ContentTypeJson)).Assert(mockT).
		JsonUnmarshallsTo([]Obj{}).
		Test()
	assert.True(t, mockT.Failed())
}

func TestExpectJsonMatches(t *testing.T) {
	mockT := new(testing.T)
	Call(setBody(`[{"id": 1}]`, ContentTypeJson)).Assert(mockT).
		JsonMatches(func(t *testing.T, list []Obj) {
			if len(list) != 1 {
				t.Errorf("Expected length 0")
			}
			if len(list) < 1 || list[0].Id != 1 {
				t.Errorf("Expected list[0].id=1")
			}
		}).Test()

	assert.False(t, mockT.Failed())
}

func TestExpectJsonMatchesCantUnmarshall(t *testing.T) {
	mockT := new(testing.T)
	Call(setBody(`[{"id": 1}]`, ContentTypeJson)).Assert(mockT).
		JsonMatches(func(t *testing.T, obj Obj) {}).Test()
	assert.True(t, mockT.Failed())
}

func TestExpectJsonMatchesFails(t *testing.T) {
	mockT := new(testing.T)
	Call(setBody(`[{"id": 1}]`, ContentTypeJson)).Assert(mockT).
		// TODO allow to use pointers also JsonMatches(func(t *testing.T, list *[]Obj) {
		JsonMatches(func(t *testing.T, list []Obj) {
			t.Errorf("Fail because something didn't meet your expectations")
		}).Test()
	assert.True(t, mockT.Failed())
}

func TestExpectJsonMatchesWrongFunc(t *testing.T) {
	type test struct {
		name     string
		function interface{}
	}
	for _, tt := range []test{
		{"Empty func", func() {}},

		{"T should be the first arg", func(Obj, Obj) {}},
		{"T should be a pointer", func(testing.T, Obj) {}},
		{"Too many params", func(*testing.T, Obj, Obj) {}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			mockT := new(testing.T)
			Call(setBody(`[{"id": 1}]`, ContentTypeJson)).Assert(mockT).
				JsonMatches(tt.function).Test()
			assert.True(t, mockT.Failed())
		})
	}
}

// TODO ConformsToFails on unexported declaration

// TODO test json requests if ContentType not set properly
