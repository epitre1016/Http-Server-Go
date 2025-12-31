package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayer(t *testing.T) {
	t.Run("request player score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/TestPlayer", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
