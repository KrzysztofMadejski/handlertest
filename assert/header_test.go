package assert_test

import (
	"github.com/krzysztofmadejski/handlertest"
	"net/http"
	"testing"
)

func setHeader(key string, val string) func(w http.ResponseWriter, r *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(key, val)
	}
	return handler
}

func TestExpectsHeader(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(setHeader("Allow-Origin", "*")).Assert(mockT).
		Header("Allow-Origin", "*")
	if mockT.Failed() {
		t.Errorf("Expected assertion to pass")
	}
}

func TestExpectsHeaderFails(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(emptyHandler).Assert(mockT).
		Header("Allow-Origin", "*")

	if !mockT.Failed() {
		t.Errorf("Assertion should fail when Header is not set")
	}
}

func TestExpectsMissingHeader(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(emptyHandler).Assert(mockT).
		HeaderMissing("Allow-Origin")

	if mockT.Failed() {
		t.Errorf("Expected assertion to pass")
	}
}

func TestExpectsMissingHeaderFails(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(setHeader("Allow-Origin", "*")).Assert(mockT).
		HeaderMissing("Allow-Origin")

	// TODO how to test message returned by our framework?
	if !mockT.Failed() {
		t.Errorf("Expected assertion to fail")
	}
}

func TestExpectsDifferentHeaderValue(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(setHeader("Allow-Origin", "http://example.com")).Assert(mockT).
		Header("Allow-Origin", "*")

	if !mockT.Failed() {
		t.Errorf("Expected assertion to fail")
	}
}

func TestExpectsContentType(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(setHeader("Content-Type", handlertest.ContentTypeJSON)).Assert(mockT).
		ContentType(handlertest.ContentTypeJSON)

	if mockT.Failed() {
		t.Errorf("Expected assertion to pass")
	}
}

func TestExpectsContentTypeFails(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(emptyHandler).Assert(mockT).
		ContentType(handlertest.ContentTypeJSON)

	if !mockT.Failed() {
		t.Errorf("Expected assertion to fail")
	}
}

func TestExpectsDifferentContentType(t *testing.T) {
	mockT := new(testing.T)
	handlertest.Call(setHeader("Content-Type", handlertest.ContentTypeFormURLEncoded)).Assert(mockT).
		ContentType(handlertest.ContentTypeJSON)

	if !mockT.Failed() {
		t.Errorf("Expected assertion to fail")
	}
}
