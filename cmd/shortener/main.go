package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Ron-ttt/xxxxxxxx/internal/app/handlers"
	"github.com/Ron-ttt/xxxxxxxx/internal/app/middleware"
	"github.com/gorilla/mux"
)

type userurl struct {
	user string
	url  string
}

func main() {

	hw := handlers.Init()
	r := mux.NewRouter()
	r.Use(middleware.Logger1, middleware.GzipMiddleware, middleware.AuthMiddleware)
	r.HandleFunc("/api/user/urls", hw.ListUserURLs).Methods(http.MethodGet)
	r.HandleFunc("/api/user/urls", hw.DeleteURL).Methods(http.MethodDelete)
	r.HandleFunc("/ping", hw.BD).Methods(http.MethodGet)
	r.HandleFunc("/", hw.IndexPage).Methods(http.MethodPost)
	r.HandleFunc("/{id}", hw.Redirect).Methods(http.MethodGet)
	r.HandleFunc("/api/shorten", hw.IndexPageJ).Methods(http.MethodPost)
	r.HandleFunc("/api/shorten/batch", hw.IndexPageM).Methods(http.MethodPost)

	log.Println("server is running")
	err := http.ListenAndServe(hw.Localhost, r)

	if err != nil {
		panic(err)
	}

	url := make(chan userurl, 100)
	for i := 0; i < 10; i++ {
		go del(hw, url)
	}
}
func del(s handlers.handlerWrapper, url <-chan string) {
	row := s.conn.QueryRow(context.Background(), "SELECT users FROM hui WHERE shorturl=$1", u)
	var name string
	err := row.Scan(&name)
	if err != nil {
		log.Println(err)
	}
	if name == user {
		_, err1 := s.conn.Exec(context.Background(), "UPDATE hui SET isDeleted=TRUE WHERE shorturl=$1", u)
		if err1 != nil {
			log.Println(err1)
		}
	}

}
