package handlers

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
	NewRequest(setHeader("Allow-Origin", "*")).Assert().
		Header("Allow-Origin", "*").
		Test(mockT)
	assert.False(t, mockT.Failed())
}

func TestExpectsHeaderFails(t *testing.T) {
	mockT := new(testing.T)
	NewRequest(emptyHandler).Assert().
		Header("Allow-Origin", "*").
		Test(mockT)

	assert.True(t, mockT.Failed(), "Assertion should fail when Header is not set")
}

func TestExpectsMissingHeader(t *testing.T) {
	mockT := new(testing.T)
	NewRequest(emptyHandler).Assert().
		HeaderMissing("Allow-Origin").
		Test(mockT)

	assert.False(t, mockT.Failed())
}

func TestExpectsMissingHeaderFails(t *testing.T) {
	mockT := new(testing.T)
	NewRequest(setHeader("Allow-Origin", "*")).Assert().
		HeaderMissing("Allow-Origin").
		Test(mockT)

	// TODO how to test message returned by our framework?
	assert.True(t, mockT.Failed())
}

func TestExpectsDifferentHeaderValue(t *testing.T) {
	mockT := new(testing.T)
	NewRequest(setHeader("Allow-Origin", "http://example.com")).Assert().
		Header("Allow-Origin", "*").
		Test(mockT)

	assert.True(t, mockT.Failed())
}

func TestExpectsContentType(t *testing.T) {
	mockT := new(testing.T)
	NewRequest(setHeader("Content-Type", ContentTypeJson)).Assert().
		ContentType(ContentTypeJson).
		Test(mockT)

	assert.False(t, mockT.Failed())
}

func TestExpectsContentTypeFails(t *testing.T) {
	mockT := new(testing.T)
	NewRequest(emptyHandler).Assert().
		ContentType(ContentTypeJson).
		Test(mockT)

	assert.True(t, mockT.Failed())
}

func TestExpectsDifferentContentType(t *testing.T) {
	mockT := new(testing.T)
	NewRequest(setHeader("Content-Type", ContentTypeFormUrlEncoded)).Assert().
		ContentType(ContentTypeJson).
		Test(mockT)

	assert.True(t, mockT.Failed())
}
