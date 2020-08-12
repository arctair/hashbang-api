package main

import (
	"encoding/json"
	"net/http"
)

type controller struct{}

func (c *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal([]Post{})
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}
