package handlers

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/Ron-ttt/xxxxxxxx/internal/app/storage"
	"github.com/gorilla/mux"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

const localhost = "http://localhost:8080/"

func randString(n int) string {
	//rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func IndexPage(res http.ResponseWriter, req *http.Request) {
	originalURL, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(originalURL)
	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	length := 6 // Укажите длину строки
	rez1 := randString(length)
	rez := localhost + rez1
	storage := storage.NewMapStorage()
	storage.Add(rez1, string(originalURL))
	res.Write([]byte(rez))

}

func Redirect(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id := params["id"]

	storage := storage.NewMapStorage()
	originalURL, ok := storage.Get(id)
	if ok != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Set("Location", originalURL)
	res.WriteHeader(http.StatusTemporaryRedirect)

}
