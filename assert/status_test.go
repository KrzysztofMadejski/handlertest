package assert_test

import (
	"github.com/krzysztofmadejski/handlertest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestExpectsStatus(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	mockT1 := new(testing.T)
	handlertest.Call(handler).Assert(mockT1).Status(http.StatusOK)
	assert.False(t, mockT1.Failed())

	mockT2 := new(testing.T)
	handlertest.Call(handler).Assert(mockT2).Status(http.StatusAccepted)
	if !mockT2.Failed() {
		t.Errorf("Status assertion should fail")
	}
}
