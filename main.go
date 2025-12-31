package main

import (
	"log"
	"net/http"
)

type InMemeoryStore struct{}

func (i *InMemeoryStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&InMemeoryStore{}}
	log.Fatal(http.ListenAndServe(":5050", server))
}
