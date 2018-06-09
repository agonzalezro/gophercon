package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestMainHandler(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mainHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != message {
		t.Errorf("Expecting: %s, got: %s", message, string(body))
	}
}
