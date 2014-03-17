package main

import (
	"github.com/codegangsta/martini"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_AddHeaders(t *testing.T) {
	// Set up
	recorder := httptest.NewRecorder()

	m := martini.New()
	m.Use(addHeaders())

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
	}

	m.ServeHTTP(recorder, r)

	// Make our assertions
	_, ok := recorder.HeaderMap[HeaderContentType]
	if !ok {
		t.Error(HeaderContentType + " not present")
	}

	ct := recorder.Header().Get(HeaderContentType)
	if !strings.EqualFold(ct, "application/json; charset=utf-8") {
		t.Error(HeaderContentType + " is not 'application/json; charset=utf-8'")
	}
}
