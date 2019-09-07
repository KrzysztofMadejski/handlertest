package handlers

import (
	"net/http"
	"testing"
)

func TestImports(t *testing.T) {
	handler := func(http.ResponseWriter, *http.Request) {}

	NewRequest(handler).Assert().Test(t)
}
