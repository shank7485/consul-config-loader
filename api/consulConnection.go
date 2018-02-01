package api

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"io/ioutil"
)

var client = &http.Client{}

func (kvStruct *KeyValue) WriteKVsToConsul() error {
	for key, value := range kvStruct.kv {
		// Add logic to add KVs to Consul
		//requestPUT("localhost", key, value)
		fmt.Println("key:", key, "value", value)
	}
	fmt.Println("Wrote KVs to Consul")
	return nil
}

func requestPUT(url string, key string, value string){
	request, err := http.NewRequest("PUT", url, strings.NewReader(key + ":" + value)) // Questionable
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(contents))
	}

}

func requestGET(url string, key string){
	// TODO
}
