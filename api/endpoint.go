package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"errors"
)

type ResponseStruct struct {
	Response string `json:"response"`
}

type LoadStruct struct {
	Type *TypeStruct `json:"type"`
}

type TypeStruct struct {
	FilePath string `json:"file_path"`
}

var kvStruct = &KeyValue{kv: make(map[string]string)}

func HandlePOST(w http.ResponseWriter, r *http.Request) {

	var body LoadStruct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)

	// Check empty body logic. Its not working.
	if err != nil {
		http.Error(w, "Empty body.", 400)
		return
	}

	err = readBody(body)

	if err != nil {
		req := ResponseStruct{Response: string(err.Error())}
		json.NewEncoder(w).Encode(req)
	} else {
		req := ResponseStruct{Response: "Configuration read and default Key Values loaded to Consul"}
		json.NewEncoder(w).Encode(&req)
	}
}

func HandleGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := GetKVFromConsul(key)

	if err != nil {
		req := ResponseStruct{Response: string(err.Error())}
		json.NewEncoder(w).Encode(req)
	} else {
		// Rethink if this is way to do.
		req := ResponseStruct{Response: "key: " + key + ", value: " + value}
		json.NewEncoder(w).Encode(req)
	}
}

func HandleGETS(w http.ResponseWriter, r *http.Request) {
	var out string

	values, err := GetKVsFromConsul()

	if err != nil {
		req := ResponseStruct{Response: string(err.Error())}
		json.NewEncoder(w).Encode(req)
	} else {
		// Rethink if this is way to do.
		for _, value := range values {
			out += "\"" + value + "\" "
		}
		req := ResponseStruct{Response: out}
		json.NewEncoder(w).Encode(req)
	}
}

func readBody(body LoadStruct) error {
	if body.Type.FilePath == "" {
		return errors.New("file_path not set")
	} else if body.Type.FilePath == "default" {
		kvStruct.FileReader("default")
		err := kvStruct.WriteKVsToConsul()
		if err != nil {
			return err
		}
		return nil
	} else {
		kvStruct.FileReader(body.Type.FilePath)
		err := kvStruct.WriteKVsToConsul()
		if err != nil {
			return err
		}
		return nil
	}
}
