package handlers

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func (a *Assert) JsonBody(expectedContent string) *Assert {
	a.body = func(t *testing.T, body []byte) {
		assert.Equal(t, expectedContent, string(body)) // TODO oneliner
	}
	return a
}

func (a *Assert) Body(test func(*testing.T, []byte)) *Assert {
	a.body = test
	return a
}

func (a *Assert) JsonTypeOf(obj interface{}) *Assert {
	return a.JsonConformsTo(obj, nil)
}

func (a *Assert) JsonConformsTo(obj interface{}, test func(*testing.T, interface{})) *Assert {
	a.body = func(t *testing.T, body []byte) {
		objPtr := reflect.New(reflect.TypeOf(obj))
		if err := json.Unmarshal(body, objPtr.Interface()); err != nil {
			t.Errorf("Could not unmarshall json body to %T", obj)
		}
		if test != nil {
			test(t, objPtr.Elem().Interface())
		}
	}

	return a
}
