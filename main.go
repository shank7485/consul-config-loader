package main

import (
	"github.com/gorilla/mux"
	"github.com/shank7485/consul-default-kv-microservice/api"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", api.LoadDefaultConfigs).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
