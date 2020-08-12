package main

import (
	"io"
	"net/http"
)

type controller struct{}

func (c *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}
