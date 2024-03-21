package main

import (
	"log"
	"net/http"

	"github.com/Ron-ttt/xxxxxxxx/internal/app/handlers"
	"github.com/gorilla/mux"
)

func main() {
	hw := handlers.Init()

	r := mux.NewRouter()
	r.HandleFunc("/", hw.IndexPage).Methods(http.MethodPost)
	r.HandleFunc("/{id}", hw.Redirect).Methods(http.MethodGet)

	log.Println("server is running")
	err := http.ListenAndServe(handlers.Localhost, r)

	if err != nil {
		panic(err)
	}
}
