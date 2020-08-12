package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController(t *testing.T) {
	controller := &controller{}
	t.Run("GET / greets me", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		controller.ServeHTTP(response, request)

		got := response.Body.String()
		want := "Hello world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
