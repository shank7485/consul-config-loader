package api

import (
	"encoding/json"
	//"fmt"
	"fmt"
	"net/http"
)

type RequestStuct struct {
	RequestString string `json:"word"`
}

var kvStruct = &KeyValue{kv: make(map[string]string)}

func LoadDefaultConfigs(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if fileTypes, ok := query["type"]; ok {
		for _, file := range fileTypes {
			if file == "yaml" {
				kvStruct.FileReader("yaml")
			}
			if file == "properties" {
				kvStruct.FileReader("properties")
			}
		}

		nil := kvStruct.WriteKVsToConsul()
		if nil != nil {
			fmt.Println("Configuration Loaded.")
		}

		req := RequestStuct{RequestString: "Configuration read and default Key Values loaded to Consul"}
		json.NewEncoder(w).Encode(req)
	}

	if getQueries, ok := query["get"]; ok {
		for _, getReq := range getQueries {
			value := GetKVsFromConsul(getReq)
			req := RequestStuct{RequestString: value}
			json.NewEncoder(w).Encode(req)
		}
	}

}
