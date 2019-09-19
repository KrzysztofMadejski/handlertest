package handlertest

import (
	"net/http"
	"testing"
)

func TestImports(t *testing.T) {
	handler := func(http.ResponseWriter, *http.Request) {}

	Call(handler).Assert(t).Test()
}

// TODO find and replace
var emptyHandler = func(w http.ResponseWriter, r *http.Request) {}
