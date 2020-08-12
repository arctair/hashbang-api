package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestController(t *testing.T) {
	controller := &controller{}
	t.Run("GET / returns empty list", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		controller.ServeHTTP(response, request)

		var got []Post

		if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
			t.Fatal(err)
		}
		want := []Post{
			Post{
				[]string{"https://unsplash.com/photos/aADnEWskMII/download?w=640"},
				[]string{"#tdd"},
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
