package handlers

import "testing"

func handle(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
