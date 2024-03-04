package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var rez string
var rez1 string
var originalURL []byte
var m = make(map[string]string)
var localhost = "http://localhost:8080/"

func randString(n int) string {
	//rand.NewSource(time.Now().UnixNano())
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
	//res.Write([]byte("кастрат"))
	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	length := 6 // Укажите длину строки
	rez1 = randString(length)
	rez = localhost + rez1
	m[rez1] = string(originalURL)
	//res.Write([]byte(rez))

}

func redirect(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	originalURL, ok := m[id]
	if !ok {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.Header().Set("Location", originalURL)
	res.WriteHeader(http.StatusTemporaryRedirect)
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
