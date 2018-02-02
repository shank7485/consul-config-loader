package api

import (
	"fmt"
	"sync"
)

type KeyValue struct {
	sync.RWMutex
	kv map[string]string
}

func (kvStruct *KeyValue) FileReader(typ string) {
	if typ == "yaml" {
		defer kvStruct.Unlock()
		kvStruct.Lock()

		yamlValues := YAMLtoKV()
		for key, value := range yamlValues {
			kvStruct.kv[key] = value
		}
	}

	if typ == "properties" {
		defer kvStruct.Unlock()
		kvStruct.Lock()

		yamlValues := PropertiestoKV()
		for key, value := range yamlValues {
			kvStruct.kv[key] = value
		}
	}
}

func YAMLtoKV() map[string]string {
	kvs := make(map[string]string)
	fmt.Println("Read YAML.")
	// Add read YAML Logic here.
	kvs["DB_IP"] = "192.168.0.1"
	return kvs
}

func PropertiestoKV() map[string]string {
	kvs := make(map[string]string)
	fmt.Println("Read Properties.")
	// Add read Properties file Logic here.
	kvs["REDIS_IP"] = "192.168.0.2"
	return kvs
}
