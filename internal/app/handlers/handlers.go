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

// var localhost = "http://" + Localhost + "/"

func Init() handlerWrapper {
	var localhost, baseURL = config.Flags()
	return handlerWrapper{storageInterface: storage.NewMapStorage(), Localhost: localhost, baseURL: baseURL + "/"}

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
	fmt.Print(originalURL)
	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)

	//if len(baseURL) > 22 {
	//	hw.storageInterface.Add(baseURL[22:], string(originalURL)) //хуйня работает только с 4х значным портом
	//	res.Write([]byte(baseURL))
	//} else {
	length := 6 // Укажите длину строки
	randomString := utils.RandString(length)
	rez := hw.baseURL + randomString
	hw.storageInterface.Add(randomString, string(originalURL))
	res.Write([]byte(rez))
	//}
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
