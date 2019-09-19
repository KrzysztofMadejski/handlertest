package assert

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"reflect"
	"testing"
)

func (a *Assert) JsonBody(expectedContent string) *Assert {
	return a.Body(func(t *testing.T, body []byte) {
		assert.Equal(t, expectedContent, string(body)) // TODO oneliner response + diff
	})
}

func (a *Assert) Body(test func(*testing.T, []byte)) *Assert {
	body, err := ioutil.ReadAll(a.R.Body)
	if err != nil {
		a.T.Errorf("Could not read response body: %v", err)
	}

	test(a.T, body)

	return a
}

func (a *Assert) JsonUnmarshallsTo(obj interface{}) *Assert {
	return a.Body(func(t *testing.T, body []byte) {

		objPtr := reflect.New(reflect.TypeOf(obj))
		if err := json.Unmarshal(body, objPtr.Interface()); err != nil {
			t.Errorf("Could not unmarshall json body to %T", obj)
		}
	})
}

func (a *Assert) JsonMatches(test interface{}) *Assert {
	return a.Body(func(t *testing.T, body []byte) {
		var nilT *testing.T

		f := reflect.ValueOf(test)
		if f.Kind() != reflect.Func || f.Type().NumIn() != 2 || f.Type().In(0) != reflect.TypeOf(nilT) {
			t.Errorf("Function passed to JsonConformsTo should have func(*testing.T, expectedType) ..")
			return
		}

		objPtr := reflect.New(f.Type().In(1))
		if err := json.Unmarshal(body, objPtr.Interface()); err != nil {
			t.Errorf("Could not unmarshall json body to %T", objPtr.Elem().Interface())
		}

		f.Call([]reflect.Value{
			reflect.ValueOf(t),
			objPtr.Elem(),
		})
	})
}
