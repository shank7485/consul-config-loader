package api

import (
	"errors"
	"sync"
)

type KeyValue struct {
	sync.RWMutex
	kv map[string]string
}

type ResponseStringStruct struct {
	Response string `json:"response"`
}

type ResponseGETStruct struct {
	Response map[string]string `json:"response"`
}

type ResponseGETSStruct struct {
	Response []string `json:"response"`
}

type LoadStruct struct {
	Type *TypeStruct `json:"type"`
}

type TypeStruct struct {
	FilePath string `json:"file_path"`
}

var KVStruct = &KeyValue{kv: make(map[string]string)}

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

func readBody(body LoadStruct) error {
	if body.Type == nil {
		return errors.New("Type not set. Recheck POST data.")
	}
	if body.Type.FilePath == "" {
		return errors.New("file_path not set")
	} else if body.Type.FilePath == "default" {
		KVStruct.FileReader("default")
		err := KVStruct.WriteKVsToConsul()
		if err != nil {
			return err
		}
		return nil
	} else {
		KVStruct.FileReader(body.Type.FilePath)
		err := KVStruct.WriteKVsToConsul()
		if err != nil {
			return err
		}
		return nil
	}
}
