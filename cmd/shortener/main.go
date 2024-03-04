package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var M map[string]string

const localhost = "http://localhost:8080/"

func randString(n int) string {
	//rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func indexPage(res http.ResponseWriter, req *http.Request) {

	originalURL, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(originalURL)
	//res.Write([]byte("кастрат"))
	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	length := 6 // Укажите длину строки
	rez1 := randString(length)
	rez := localhost + rez1
	M[rez1] = string(originalURL)
	res.Write([]byte(rez))

}

func redirect(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id := params["id"]

	originalURL, ok := M[id]
	if !ok {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Set("Location", originalURL)
	res.WriteHeader(http.StatusTemporaryRedirect)

}

func main() {
	M = make(map[string]string)
	r := mux.NewRouter()
	r.HandleFunc("/", indexPage).Methods(http.MethodPost)
	r.HandleFunc("/{id}", redirect).Methods(http.MethodGet)
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		panic(err)
	}
}
