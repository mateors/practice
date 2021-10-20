package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndxHandler(t *testing.T) {

	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	indxHandler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status: 200, got: %d", recorder.Code)
	}

}
