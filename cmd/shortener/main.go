package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var rez string
var originalURL []byte

func randString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func indexPage(res http.ResponseWriter, req *http.Request) {
	var err error
	originalURL, err = io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(originalURL)
	res.Write([]byte("кастрат"))
	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	length := 6 // Укажите длину строки
	rez = "http://localhost:8080/" + (randString(length))
	res.Write([]byte(rez))

}

func redirect(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("location", string(originalURL))
	res.WriteHeader(http.StatusTemporaryRedirect)
	res.Write(originalURL)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexPage).Methods(http.MethodPost)
	r.HandleFunc(`/{id}`, redirect).Methods(http.MethodGet)
	err := http.ListenAndServe(`:8080`, r)
	if err != nil {
		panic(err)
	}
}
