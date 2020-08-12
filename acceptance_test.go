package main_test

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"sync"
	"testing"

	. "arctair.com/hashbang"
)

func assertNotError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func TestAcceptance(t *testing.T) {
	serverExit := &sync.WaitGroup{}
	serverExit.Add(1)
	server := StartHTTPServer(serverExit)

	defer func() {
		if err := server.Shutdown(context.TODO()); err != nil {
			panic(err)
		}

		serverExit.Wait()
	}()

	response, err := http.Get("http://localhost:5000/")
	assertNotError(t, err)

	var got []Post
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&got)
	assertNotError(t, err)

	want := []Post{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
