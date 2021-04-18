package main

import (
	"net/http"

	"github.com/gorilla/pat"
)

func hostAPI() {
	mux := pat.New()
	srv := http.Server{
		Addr:    ":3003",
		Handler: mux,
	}

	//Normal listing
	mux.Get("/dunkirk", http.HandlerFunc(shouldIHaveMyLightsOn))

	go srv.ListenAndServe()
}
