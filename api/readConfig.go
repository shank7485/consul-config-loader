package api

import (
	"sync"
)

type KeyValue struct {
	sync.RWMutex
	kv map[string]string
}

func (kvStruct *KeyValue) FileReader(typ string) {
	if typ == "properties" {
		defer kvStruct.Unlock()
		kvStruct.Lock()

		propertiesValues := PropertiesToKV()
		for key, value := range propertiesValues {
			kvStruct.kv[key] = value
		}
	}
}
