package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Ron-ttt/xxxxxxxx/internal/app/storage"
	"github.com/Ron-ttt/xxxxxxxx/internal/app/utils"

	"github.com/Ron-ttt/xxxxxxxx/internal/app/config"

	"github.com/gorilla/mux"
)

// var localhost = "http://" + Localhost + "/"

func Init() handlerWrapper {
	var localhost, baseURL = config.Flags()
	return handlerWrapper{storageInterface: storage.NewMapStorage(), Localhost: localhost, baseURL: baseURL + "/"}

}
func MInit() handlerWrapper {
	return handlerWrapper{storageInterface: storage.NewMockStorage(), Localhost: "localhost:8080", baseURL: "http://localhost:8080/"}
}

type handlerWrapper struct {
	storageInterface storage.Storage
	Localhost        string
	baseURL          string
}

func (hw handlerWrapper) IndexPage(res http.ResponseWriter, req *http.Request) { // post
	originalURL, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	_, err1 := url.Parse(string(originalURL))
	if err1 != nil {
		panic(err1)
	}
	fmt.Print(originalURL)
	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)

	length := 6 // Укажите длину строки
	randomString := utils.RandString(length)
	rez := hw.baseURL + randomString
	hw.storageInterface.Add(randomString, string(originalURL))
	res.Write([]byte(rez))

}

func (hw handlerWrapper) Redirect(res http.ResponseWriter, req *http.Request) { //get
	params := mux.Vars(req)

	id := params["id"]

	originalURL, ok := hw.storageInterface.Get(id)
	if ok != nil {
		http.Error(res, "not found", http.StatusBadRequest)
		return
	}

	res.Header().Set("Location", originalURL)
	res.WriteHeader(http.StatusTemporaryRedirect)

}
