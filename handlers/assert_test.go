package handlers

import (
	"net/http"
	"testing"
)

func TestHandlerIsCalled(t *testing.T) {
	called := false
	handler := func(w http.ResponseWriter, r *http.Request) {
		called = true
	}

	NewRequest(handler).Assert().Test(new(testing.T))

	if !called {
		t.Errorf("Handler was not called")
	}
}

func TestEmptyTest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {}

	mockT1 := new(testing.T)
	NewRequest(handler).Assert().Test(mockT1)
	if mockT1.Failed() {
		t.Errorf("Empty test should not fail")
	}
}

// TODO test for TestRun
// it seems that t creates a child t
// t.Run("POST", NewRequest(expectMethod("POST")).POST("/jobs").Assert().TestRun())
