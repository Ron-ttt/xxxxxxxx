package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/Ron-ttt/xxxxxxxx/internal/app/config"
	"github.com/Ron-ttt/xxxxxxxx/internal/app/storage"
	"github.com/Ron-ttt/xxxxxxxx/internal/app/utils"

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
	localhost, baseURL, storageType, dbAdress := config.Flags()
	if dbAdress != "" {
		dBStorage, err := storage.NewDBStorage(dbAdress)
		if err == nil {
			return handlerWrapper{storageInterface: dBStorage, Localhost: localhost, baseURL: baseURL + "/"}
		}
	}
	if storageType != "" {
		fileStorage, err := storage.NewFileStorage(storageType)
		if err == nil {
			return handlerWrapper{storageInterface: fileStorage, Localhost: localhost, baseURL: baseURL + "/"}
		}
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
		return
	}
	_, err1 := url.ParseRequestURI(string(originalURL))
	if err1 != nil {
		http.Error(res, "invalid url", http.StatusBadRequest)
		return
	}
	length := 6 // Укажите длину строки
	randomString := utils.RandString(length)
	rez := hw.baseURL + randomString
	err = hw.storageInterface.Add(randomString, string(originalURL))
	if err != nil {
		http.Error(res, "ошибка эдд", http.StatusBadRequest)
		return
	}
	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(rez))
}

func (hw handlerWrapper) IndexPageM(res http.ResponseWriter, req *http.Request) { // post
	var body []storage.URLRegistryM
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	l := len(body)
	for i := 0; i < l; i++ {
		_, err1 := url.ParseRequestURI(body[i].OriginalUrl)
		if err1 != nil {
			http.Error(res, "invalid url", http.StatusBadRequest)
			return
		}
	}
	res.Header().Set("content-type", "application/json")
	res.WriteHeader(http.StatusCreated)
	var listshort []string
	var rez []storage.URLRegistryMRes
	length := 6 // Укажите длину строки
	for i := 0; i < l; i++ {
		randomString := utils.RandString(length)
		listshort = append(listshort, randomString)
		rez[i].Id = body[i].Id
		rez[i].ShortUrl = hw.baseURL + randomString
	}
	hw.storageInterface.AddM(body, listshort)

	if err := json.NewEncoder(res).Encode(rez); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
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
		return
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

func (hw handlerWrapper) BD(res http.ResponseWriter, req *http.Request) {
	err := hw.storageInterface.Ping()
	if err != nil {
		http.Error(res, "нет бд", http.StatusBadRequest)
	} else {
		res.WriteHeader(http.StatusOK)
	}
}
