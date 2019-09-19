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

func (a *Assert) Status(statusCode int) *Assert {
	if a.R.StatusCode != statusCode {
		a.T.Errorf("Expected statusCode %d, got %d", statusCode, a.R.StatusCode)
	}

	return a
}

func (a *Assert) Header(key string, value string) *Assert {
	values := a.R.Header[key]
	if values == nil || len(values) == 0 {
		a.T.Errorf("Expected header %s to be set, it is not", key)

	} else if got := a.R.Header.Get(key); got != value {
		a.T.Errorf("Expected header %s to be set to '%s', got '%s'", key, value, got)
	}

	return a
}

func (a *Assert) HeaderMissing(key string) *Assert {
	value := a.R.Header.Get(key)
	if value != "" {
		a.T.Errorf("Expected header %s to be empty, got '%s'", key, value)
	}

	return a
}

func (a *Assert) ContentType(contentType string) *Assert {
	return a.Header("Content-Type", contentType)
}

func (a *Assert) Custom(customTest func(t *testing.T, response *http.Response)) *Assert {
	customTest(a.T, a.R)

	return a
}
