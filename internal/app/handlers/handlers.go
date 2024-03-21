package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Ron-ttt/xxxxxxxx/internal/app/storage"
	"github.com/Ron-ttt/xxxxxxxx/internal/app/utils"

	"github.com/Ron-ttt/xxxxxxxx/internal/app/config"

	"github.com/gorilla/mux"
)

var Localhost, baseURL = config.Flags()
var localhost = "http://" + Localhost + "/"

func Init() handlerWrapper {
	return handlerWrapper{storageInterface: storage.NewMapStorage()}

}

type handlerWrapper struct {
	storageInterface storage.Storage
}

func (hw handlerWrapper) IndexPage(res http.ResponseWriter, req *http.Request) { // post
	originalURL, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(originalURL)
	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)

	if len(baseURL) > 1 {
		hw.storageInterface.Add(baseURL[22:], string(originalURL)) //хуйня работает только с 4х значным портом
		res.Write([]byte(baseURL))
	} else {
		length := 6 // Укажите длину строки
		rez1 := utils.RandString(length)
		rez := localhost + rez1
		hw.storageInterface.Add(rez1, string(originalURL))
		res.Write([]byte(rez))
	}
}

func (hw handlerWrapper) Redirect(res http.ResponseWriter, req *http.Request) { //get
	params := mux.Vars(req)

	id := params["id"]

	originalURL, ok := hw.storageInterface.Get(id)
	if ok != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Set("Location", originalURL)
	res.WriteHeader(http.StatusTemporaryRedirect)

}
