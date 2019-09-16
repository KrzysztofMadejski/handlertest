package handlers

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func setBody(content string, contentType string) func(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(content))
	}
	return handler
}

func TestExpectsBodyFunction(t *testing.T) {
	mockT := new(testing.T)
	NewRequest(setBody(`[{"id": 201809}]`, ContentTypeJson)).Assert().
		Body(func(t *testing.T, body []byte) {
			// don't raise error on mockT
		}).Test(mockT)
	assert.False(t, mockT.Failed())
}

type Obj struct {
	Id int `json:"id"`
}

func TestExpectsBodyFunctionFails(t *testing.T) {
	mockT := new(testing.T)
	NewRequest(setBody(`[{"id": 201809}]`, ContentTypeJson)).Assert().
		Body(func(t *testing.T, body []byte) {
			var o Obj
			if err := json.Unmarshal(body, o); err != nil {
				t.Errorf("Could not unmarshall body")
				return
			}
			if o.Id > 201807 {
				mockT.Errorf("Expected Id to be something it wasn't")
			}
		}).Test(mockT)
	assert.True(t, mockT.Failed())
}

// TODO charset: utf-8 in Content-Type
func TestExpectsJsonBody(t *testing.T) {
	mockT := new(testing.T)
	NewRequest(setBody(`[]`, ContentTypeJson)).Assert().
		JsonBody(`[]`).
		Test(mockT)
	assert.False(t, mockT.Failed())
}

func TestExpectsJsonBodyFails(t *testing.T) {
	mockT := new(testing.T)
	NewRequest(setBody(`[]`, ContentTypeJson)).Assert().
		JsonBody(`[{"id": 1}]`).
		Test(mockT)
	assert.True(t, mockT.Failed(), "Assertion should fail when body is different")
}

// TODO test json indenting
// TODO test error message returns a nice diff during extended run

func TestExpectJsonType(t *testing.T) {
	mockT := new(testing.T)
	NewRequest(setBody(`[{"id": 1}]`, ContentTypeJson)).Assert().
		JsonTypeOf([]Obj{}).
		Test(mockT)
	assert.False(t, mockT.Failed())
}

func TestExpectJsonTypeFails(t *testing.T) {
	mockT := new(testing.T)
	NewRequest(setBody(`{"id": 1}`, ContentTypeJson)).Assert().
		JsonTypeOf([]Obj{}).
		Test(mockT)
	assert.True(t, mockT.Failed())
}

func TestExpectJsonConformsTo(t *testing.T) {
	mockT := new(testing.T)
	NewRequest(setBody(`[{"id": 1}]`, ContentTypeJson)).Assert().
		JsonConformsTo([]Obj{}, func(t *testing.T, ret interface{}) {
			list := ret.([]Obj)
			if len(list) != 1 {
				t.Errorf("Expected length 0")
			}
			if len(list) < 1 || list[0].Id != 1 {
				t.Errorf("Expected list[0].id=1")
			}
		}).Test(mockT)

	assert.False(t, mockT.Failed())
}

func TestExpectJsonConformsToFails(t *testing.T) {
	mockT := new(testing.T)
	NewRequest(setBody(`[{"id": 1}]`, ContentTypeJson)).Assert().
		JsonConformsTo([]Obj{}, func(t *testing.T, ret interface{}) {
			t.Errorf("Fail because something didn't meet your expectations")
		}).Test(mockT)
	assert.True(t, mockT.Failed())
}

// TODO ConformsToFails on unexported declaration

// TODO test json requests if ContentType not set properly
