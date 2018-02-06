package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func HandlePOST(w http.ResponseWriter, r *http.Request) {

	var body LoadStruct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)

	if err != nil {
		req := ResponseStringStruct{Response: "Empty body."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&req)
	}

	err = readBody(body)

	if err != nil {
		req := ResponseStringStruct{Response: string(err.Error())}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(req)
	} else {
		req := ResponseStringStruct{Response: "Configuration read and default Key Values loaded to Consul"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&req)
	}
}

func HandleGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := GetKVFromConsul(key)

	if err != nil {
		req := ResponseStringStruct{Response: string(err.Error())}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(req)
	} else {
		req := ResponseGETStruct{Response: map[string]string{key: value}}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(req)
	}
}

func HandleGETS(w http.ResponseWriter, r *http.Request) {

	values, err := GetKVsFromConsul()

	if err != nil {
		req := ResponseStringStruct{Response: string(err.Error())}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(req)
	} else {
		req := ResponseGETSStruct{Response: values}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(req)
	}
}
