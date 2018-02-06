package main

import (
	"github.com/gorilla/mux"
	"github.com/shank7485/consul-default-kv-microservice/api"
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
