package main

import (
	"encoding/json"
	"net/http"
)

type controller struct{}

func (c *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(
		[]Post{
			Post{
				[]string{"https://unsplash.com/photos/aADnEWskMII/download?w=640"},
				[]string{"#tdd"},
			},
		},
	)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}
