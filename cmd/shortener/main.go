package main

import (
	"log"
	"net/http"

	"github.com/Ron-ttt/xxxxxxxx/internal/app/handlers"
	"github.com/Ron-ttt/xxxxxxxx/internal/app/middleware"
	"github.com/gorilla/mux"
)

func main() {
	hw := handlers.Init()

	r := mux.NewRouter()
	r.Use(middleware.Logger1, middleware.GzipMiddleware)
	r.HandleFunc("/", hw.IndexPage).Methods(http.MethodPost)
	r.HandleFunc("/{id}", hw.Redirect).Methods(http.MethodGet)
	r.HandleFunc("/api/shorten", hw.IndexPageJ).Methods(http.MethodPost)

	log.Println("server is running")
	err := http.ListenAndServe(hw.Localhost, r)

	if err != nil {
		panic(err)
	}
}
