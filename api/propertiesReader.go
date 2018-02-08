package api

import (
	"errors"
	"github.com/magiconair/properties"
	"io/ioutil"
	"path"
	"runtime"
)

func PropertiesFilesToKV(directory string) (map[string]string, error) {
	if directory == "default" {
		kvs := make(map[string]string)

		_, filename, _, ok := runtime.Caller(0)

		if !ok {
			return nil, errors.New("No caller")
		}

		configDir := path.Dir(filename) + "/../configurations/"
		err := ReadMultipleProperties(configDir, kvs)
		if err != nil {
			return nil, err
		}
		return kvs, nil
	} else {
		// Add case if directory is not there.
		kvs := make(map[string]string)
		directory += "/"
		err := ReadMultipleProperties(directory, kvs)
		if err != nil {
			return nil, err
		}
		return kvs, nil
	}
}

func ReadProperty(path string, kvs map[string]string) {
	p := properties.MustLoadFile(path, properties.UTF8)
	for _, key := range p.Keys() {
		kvs[key] = p.MustGet(key)
	}
}

func ReadMultipleProperties(path string, kvs map[string]string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		ReadProperty(path+f.Name(), kvs)
	}

	return nil
}
