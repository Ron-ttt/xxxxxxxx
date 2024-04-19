package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/Ron-ttt/xxxxxxxx/internal/app/storage"
	"github.com/Ron-ttt/xxxxxxxx/internal/app/utils"

	"github.com/Ron-ttt/xxxxxxxx/internal/app/config"

	"github.com/gorilla/mux"
)

type URLRegistry struct {
	URL string `json:"url"`
}

type URLRegistryResult struct {
	Result string `json:"result"`
}

// var localhost = "http://" + Localhost + "/"

func Init() handlerWrapper {
	var localhost, baseURL, storageType = config.Flags()
	if storageType != "" {
		fileStorage, err := storage.NewFileStorage(storageType)
		if err != nil {
			log.Fatal("unable to create file storage")
		}
		return handlerWrapper{storageInterface: fileStorage, Localhost: localhost, baseURL: baseURL + "/"}
	}
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
		http.Error(res, "unable to read body", http.StatusBadRequest)
	}
	_, err1 := url.ParseRequestURI(string(originalURL))
	if err1 != nil {
		http.Error(res, "invalid url", http.StatusBadRequest)
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

func (hw handlerWrapper) IndexPageJ(res http.ResponseWriter, req *http.Request) { // post
	var longURL URLRegistry
	if err := json.NewDecoder(req.Body).Decode(&longURL); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	_, err1 := url.ParseRequestURI(longURL.URL)
	if err1 != nil {
		http.Error(res, "invalid url", http.StatusBadRequest)
	}

	res.Header().Set("content-type", "application/json")
	res.WriteHeader(http.StatusCreated)

	length := 6 // Укажите длину строки
	randomString := utils.RandString(length)
	var rez URLRegistryResult
	rez.Result = hw.baseURL + randomString
	hw.storageInterface.Add(randomString, string(longURL.URL))
	if err := json.NewEncoder(res).Encode(rez); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
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
