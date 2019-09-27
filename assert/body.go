package assert

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func CompactJsonb(jsonBytes []byte, t *testing.T) string {
	dst := new(bytes.Buffer)
	if err := json.Compact(dst, jsonBytes); err != nil {
		t.Error(err)
		return ""
	}

	return dst.String()
}
func CompactJson(jsonStr string, t *testing.T) string {
	return CompactJsonb([]byte(jsonStr), t)
}

func IndentJsonb(jsonBytes []byte, t *testing.T) string {
	dst := new(bytes.Buffer)
	if err := json.Indent(dst, jsonBytes, "", "\t"); err != nil {
		t.Error(err)
		return ""
	}

	return dst.String()
}
func IndentJson(jsonStr string, t *testing.T) string {
	return IndentJsonb([]byte(jsonStr), t)
}

func shouldDiff() bool {
	ret, err := strconv.ParseBool(os.Getenv("HANDLERTEST_DIFF"))
	if err != nil {
		return false // default
	}
	return ret
}

func (a *Assert) Json(expectedContent string) *Assert {
	return a.Body(func(t *testing.T, body []byte) {
		if expectedContent == "" {
			t.Errorf("Empty string is not a valid json")
			return
		}

		var actual string
		if shouldDiff() {
			if differ != nil {
				expectedContent = IndentJson(expectedContent, t)
				actual = IndentJsonb(body, t)

				if expectedContent != "" && actual != "" && expectedContent != actual {
					diff := differ(expectedContent, actual)
					t.Error("Expected JSON response -expected,+actual ", diff)
				}
				return
			}
			t.Errorf("You set HANDLERTEST_DIFF, but didn't configure Differ with handlertest.SetDiff(d Differ)")
		}

		// compact version
		expectedContent = CompactJson(expectedContent, t)
		actual = CompactJsonb(body, t)

		if expectedContent != "" && actual != "" && expectedContent != actual {
			t.Errorf("Expected JSON response '%s', but got '%s", expectedContent, actual)
		}
	}).ContentType("application/json") // TODO contenttype.JSON)
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
