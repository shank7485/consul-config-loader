package api

import (
	"encoding/json"
	"net/http"
)

type RequestStuct struct {
	RequestString string `json:"response"`
}

var kvStruct = &KeyValue{kv: make(map[string]string)}

func LoadDefaultConfigs(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if fileTypes, ok := query["type"]; ok {
		for _, file := range fileTypes {
			if file == "properties" {
				kvStruct.FileReader("properties")
			}
		}

		err := kvStruct.WriteKVsToConsul()
		if err != nil {
			req := RequestStuct{RequestString: string(err.Error())}
			json.NewEncoder(w).Encode(req)
		} else {
			req := RequestStuct{RequestString: "Configuration read and default Key Values loaded to Consul"}
			json.NewEncoder(w).Encode(req)
		}
	}

	if getQueries, ok := query["get"]; ok {
		for _, key := range getQueries {
			value, err := GetKVsFromConsul(key)
			if err != nil {
				req := RequestStuct{RequestString: string(err.Error())}
				json.NewEncoder(w).Encode(req)
			} else {
				req := RequestStuct{RequestString: "key: " + key + ", value: " + value}
				json.NewEncoder(w).Encode(req)
			}
		}
	}

}
