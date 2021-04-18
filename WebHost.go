package main

import (
	"net/http"

	"github.com/gorilla/pat"
)

func hostAPI() {
	//Chooses the port to use
	mux := pat.New()
	srv := http.Server{
		Addr:    ":3003",
		Handler: mux,
	}

	//Redirect to the right function in API.go
	mux.Get("/dunkirk", http.HandlerFunc(shouldIHaveMyLightsOn))

	go srv.ListenAndServe()
}
