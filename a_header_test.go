package handlertest

import (
	"github.com/stretchr/testify/assert"
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
	Call(setHeader("Allow-Origin", "*")).Assert(mockT).
		Header("Allow-Origin", "*").
		Test()
	assert.False(t, mockT.Failed())
}

func TestExpectsHeaderFails(t *testing.T) {
	mockT := new(testing.T)
	Call(emptyHandler).Assert(mockT).
		Header("Allow-Origin", "*").
		Test()

	assert.True(t, mockT.Failed(), "Assertion should fail when Header is not set")
}

func TestExpectsMissingHeader(t *testing.T) {
	mockT := new(testing.T)
	Call(emptyHandler).Assert(mockT).
		HeaderMissing("Allow-Origin").
		Test()

	assert.False(t, mockT.Failed())
}

func TestExpectsMissingHeaderFails(t *testing.T) {
	mockT := new(testing.T)
	Call(setHeader("Allow-Origin", "*")).Assert(mockT).
		HeaderMissing("Allow-Origin").
		Test()

	// TODO how to test message returned by our framework?
	assert.True(t, mockT.Failed())
}

func TestExpectsDifferentHeaderValue(t *testing.T) {
	mockT := new(testing.T)
	Call(setHeader("Allow-Origin", "http://example.com")).Assert(mockT).
		Header("Allow-Origin", "*").
		Test()

	assert.True(t, mockT.Failed())
}

func TestExpectsContentType(t *testing.T) {
	mockT := new(testing.T)
	Call(setHeader("Content-Type", ContentTypeJson)).Assert(mockT).
		ContentType(ContentTypeJson).
		Test()

	assert.False(t, mockT.Failed())
}

func TestExpectsContentTypeFails(t *testing.T) {
	mockT := new(testing.T)
	Call(emptyHandler).Assert(mockT).
		ContentType(ContentTypeJson).
		Test()

	assert.True(t, mockT.Failed())
}

func TestExpectsDifferentContentType(t *testing.T) {
	mockT := new(testing.T)
	Call(setHeader("Content-Type", ContentTypeFormUrlEncoded)).Assert(mockT).
		ContentType(ContentTypeJson).
		Test()

	assert.True(t, mockT.Failed())
}
