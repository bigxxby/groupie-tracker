package main

import (
	bigxxby "bigxxby/internal"
	"fmt"
	"net/http"
)

// https://www.codementor.io/@hau12a1/golang-http-serve-static-files-correctly-m55u3vz1a

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", bigxxby.MainHandler)
	mux.HandleFunc("/static/", bigxxby.StaticHandler)
	mux.HandleFunc("/artists/", bigxxby.ArtistIdHandler)

	fmt.Println("Listening and serving on http://localhost:3000")
	http.ListenAndServe(":3000", mux)
}
