package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestStuct struct {
	RequestString string `json:"word"`
}

var kvStruct = &KeyValue{kv: make(map[string]string)}

func LoadDefaultConfigs(w http.ResponseWriter, r *http.Request) {
	fileType := r.URL.Query().Get("type")

	if fileType == "yaml" {
		kvStruct.FileReader("yaml")
	}
	if fileType == "properties" {
		kvStruct.FileReader("properties")
	}

	nil := kvStruct.WriteKVsToConsul()

	if nil != nil {
		fmt.Println("Configuration Loaded.")
	}

	req := RequestStuct{RequestString: "Configuration read and default Key Values loaded to Consul"}
	json.NewEncoder(w).Encode(req)
}
