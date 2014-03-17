package main

import (
	"github.com/codegangsta/martini"
	"net/http"
)

const HeaderContentType = "Content-Type"

func addHeaders() martini.Handler {
	return func(w http.ResponseWriter, r *http.Request, c martini.Context) {
		headers := w.Header()
		headers.Set(HeaderContentType, "application/json; charset=utf-8")
	}
}
