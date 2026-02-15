// run with go test

package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestHandleHello(t *testing.T) {
	w := httptest.NewRecorder()

	handleHello(w, nil)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code, expected %v but got %v\body: %s\n", desiredCode, w.Code, w.Body.String())
	}
	expectedMessage := []byte("Hello World\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleGoodbye(t *testing.T) {
	w := httptest.NewRecorder()

	handleGoodbye(w, nil)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code, expected %v but got %v\body: %s\n", desiredCode, w.Code, w.Body.String())
	}
	expectedMessage := [] byte("GoodBye")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("Bad return got %q, expected %q", w.Body.String(), expectedMessage)
	}
}