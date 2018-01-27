package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestStuct struct {
	RequestString string `json:"word"`
}

func LoadDefaultConfigs(w http.ResponseWriter, r *http.Request) {
	FilesReader()
	WriteKVsToConsul()

	fmt.Println("Configuration Loaded.")

	req := RequestStuct{RequestString: "Configuration read and default Key Values loaded to Consul"}
	json.NewEncoder(w).Encode(req)
}
