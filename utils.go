package handlertest

import "testing"
import "net/url"

func handle(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func ValuesFromMap(values map[string]string) url.Values {
	v := url.Values{}
	for key, val := range values {
		v.Set(key, val)
	}
	return v
}
