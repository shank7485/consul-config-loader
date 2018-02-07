package main

// Make sure imports from shank7485 github are removed and made local. This is possible by creating a folder under
// $GOPATH/src/ since that is what imports pick up from. So essentially, move consul-config-loader under $GOPATH/src
// directly before pushing to Gerrit.

import (
	"github.com/gorilla/mux"
	"github.com/shank7485/consul-config-loader/api"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/loadconfigs", api.HandlePOST).Methods("POST")
	router.HandleFunc("/getconfig/{key}", api.HandleGET).Methods("GET")
	router.HandleFunc("/getconfigs", api.HandleGETS).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
