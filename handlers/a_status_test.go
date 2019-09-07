package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestStatus(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	mockT1 := new(testing.T)
	NewRequest(handler).Assert().Status(http.StatusOK).Test(mockT1)
	assert.False(t, mockT1.Failed())

	mockT2 := new(testing.T)
	NewRequest(handler).Assert().Status(http.StatusAccepted).Test(mockT2)
	if !mockT2.Failed() {t.Errorf("Status assertion should fail")}
}
