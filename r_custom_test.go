package handlertest

import (
	"net/http"
	"testing"
)

func TestCustom(t *testing.T) {
	NewRequest(expectHeader(t, "Allow-Origin", "*")).Custom(func(request *http.Request) {
		request.Header.Set("Allow-Origin", "*")
	}).Assert().Test(new(testing.T))
}
