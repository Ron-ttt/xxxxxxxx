package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Ron-ttt/xxxxxxxx/internal/app/handlers"

	"github.com/gorilla/mux"
)

func main() {
	hw := handlers.Init()
	// Определение флагов
	address := flag.String("a", "localhost:8888", "адрес запуска HTTP-сервера")
	baseURL := flag.String("b", "http://localhost:8000/qsd54gFg", "базовый адрес результирующего сокращённого URL")

	// Парсинг флагов
	flag.Parse()

	r := mux.NewRouter()

	r.HandleFunc("/", hw.IndexPage).Methods(http.MethodPost)
	r.HandleFunc("/{id}", hw.Redirect).Methods(http.MethodGet)

	log.Println("server is running")
	err := http.ListenAndServe(*address, r)

	if err != nil {
		panic(err)
	}
}
