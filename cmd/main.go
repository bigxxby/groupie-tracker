package main

import (
	bigxxby "bigxxby/internal"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", bigxxby.MainHandler)
	mux.HandleFunc("/static/", bigxxby.StaticHandler)
	mux.HandleFunc("/artists/", bigxxby.ArtistIdHandler)

	fmt.Println("Listening and serving on http://localhost:3000")
	http.ListenAndServe(":3000", mux)
}
