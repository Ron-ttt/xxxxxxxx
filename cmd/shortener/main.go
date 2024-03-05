package main

import (
	"net/http"

	"github.com/Ron-ttt/xxxxxxxx/internal/app/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.IndexPage).Methods(http.MethodPost)
	r.HandleFunc("/{id}", handlers.Redirect).Methods(http.MethodGet)
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		panic(err)
	}
}
