package api

import (
	"github.com/magiconair/properties"
	"io/ioutil"
	"path"
	"runtime"
)

func PropertiesToKV(directory string) map[string]string {
	if directory == "default" {
		kvs := make(map[string]string)

		_, filename, _, ok := runtime.Caller(0)

		if !ok {
			panic("No caller information")
		}

		configDir := path.Dir(filename) + "/../configurations/"
		ReadMultipleProperties(configDir, kvs)
		return kvs
	} else {
		// Add case if directory is not there.
		kvs := make(map[string]string)

		_, filename, _, ok := runtime.Caller(0)

		if !ok {
			panic("No caller information")
		}

		configDir := path.Dir(filename) + "/../configurations/"
		ReadMultipleProperties(configDir, kvs)
		return kvs
	}
}

func ReadProperty(path string, kvs map[string]string) {
	p := properties.MustLoadFile(path, properties.UTF8)
	for _, key := range p.Keys() {
		kvs[key] = p.MustGet(key)
	}
}

func ReadMultipleProperties(path string, kvs map[string]string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		ReadProperty(path+f.Name(), kvs)
	}
}
