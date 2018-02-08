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

func (kvStruct *KeyValue) ReadConfigs(body LoadStruct) error {
	if body.Type.FilePath == "default" {
		err := kvStruct.FileReader("default")
		if err != nil {
			return err
		}
		return nil
	} else {
		err := kvStruct.FileReader(body.Type.FilePath)
		if err != nil {
			return err
		}
		return nil
	}
}

func (kvStruct *KeyValue) FileReader(directory string) error {
	defer kvStruct.Unlock()

	kvStruct.Lock()

	if directory == "default" {
		propertiesValues, err := PropertiesFilesToKV("default")
		if err != nil {
			return err
		}
		for key, value := range propertiesValues {
			kvStruct.kv[key] = value
		}
		return nil
	} else {
		propertiesValues, err := PropertiesFilesToKV(directory)
		if err != nil {
			return err
		}
		for key, value := range propertiesValues {
			kvStruct.kv[key] = value
		}
		return nil
	}
}

func ValidateBody(body LoadStruct) error {
	if body.Type == nil {
		return errors.New("Type not set. Recheck POST data.")
	} else if body.Type.FilePath == "" {
		return errors.New("file_path not set")
	} else {
		return nil
	}
}
