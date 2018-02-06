package api

import (
	"sync"
)

type KeyValue struct {
	sync.RWMutex
	kv map[string]string
}

func (kvStruct *KeyValue) FileReader(typ string) {
	defer kvStruct.Unlock()

	kvStruct.Lock()

	if typ == "default" {
		propertiesValues := PropertiesToKV("default")
		for key, value := range propertiesValues {
			kvStruct.kv[key] = value
		}
	} else {
		// Pass directory.
		propertiesValues := PropertiesToKV("default")
		for key, value := range propertiesValues {
			kvStruct.kv[key] = value
		}
	}
}
