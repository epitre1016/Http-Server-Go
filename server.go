package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(response http.ResponseWriter, request *http.Request) {

	player := strings.TrimPrefix(request.URL.Path, "/players/")

	fmt.Fprint(response, GetPlayerScore(player))
}

func GetPlayerScore(name string) string {

	if name == "Pepe" {
		return "20"
	}

	if name == "Charlie" {
		return "10"
	}

	return ""
}
