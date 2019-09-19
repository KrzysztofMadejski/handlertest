package assert_test

import (
	"github.com/krzysztofmadejski/handlertest"
	"net/http"
	"testing"
)

func TestHandlerIsCalled(t *testing.T) {
	called := false
	handler := func(w http.ResponseWriter, r *http.Request) {
		called = true
	}

	handlertest.Call(handler).Assert(new(testing.T))

	if !called {
		t.Errorf("Handler was not called")
	}
}

func TestEmptyTest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {}

	mockT1 := new(testing.T)
	handlertest.Call(handler).Assert(mockT1)
	if mockT1.Failed() {
		t.Errorf("Empty test should not fail")
	}
}

// TODO test for TestRun
// it seems that T creates a child T
// T.Run("POST", Call(expectMethod("POST")).POST("/jobs").Assert(mockT).TestRun())

func TestCustomAssertions(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {}
	mockT := new(testing.T)
	handlertest.Call(handler).Assert(mockT).Custom(func(t *testing.T, response *http.Response) {
		t.Error("Fail")
	})
	if !mockT.Failed() {
		t.Errorf("Expected assertion to fail")
	}
}
