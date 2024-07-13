package main

import (
	"log"
	"net/http"

	"github.com/juntakoman123/go_basic_study/server"
)

func main() {

	store := server.NewInMemoryPlayerStore()
	server := server.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}
