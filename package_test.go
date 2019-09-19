package handlertest

import (
	"net/http"
	"testing"
)

func TestImports(t *testing.T) {
	handler := func(http.ResponseWriter, *http.Request) {}

	Call(handler).Assert(t)
}
