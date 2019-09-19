package assert

import (
	"net/http"
	"testing"
)

type Assert struct {
	R *http.Response
	T *testing.T
}

func (a *Assert) TestRun() func(*testing.T) {
	return func(t *testing.T) {
		// TODO test it
	}
}

func (a *Assert) Custom(customTest func(t *testing.T, response *http.Response)) *Assert {
	customTest(a.T, a.R)

	return a
}
