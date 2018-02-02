package api

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func (kvStruct *KeyValue) WriteKVsToConsul() error {
	for key, value := range kvStruct.kv {
		// Add logic to add KVs to Consul
		requestPUT("127.0.0.1", key, value)
		fmt.Println("key:", key, "value", value)
	}
	fmt.Println("Wrote KVs to Consul")
	return nil
}

func GetKVsFromConsul(key string) string {
	return requestGET("127.0.0.1", key)
}

func requestPUT(url string, key string, value string) {
	config := api.DefaultConfig()
	config.Address = url + ":8500"
	client, err := api.NewClient(config)

	if err != nil {
		panic(err)
	}

	kv := client.KV()

	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}
}

func requestGET(url string, key string) string {
	config := api.DefaultConfig()
	config.Address = url + ":8500"
	client, err := api.NewClient(config)

	if err != nil {
		panic(err)
	}

	kv := client.KV()

	pair, _, err := kv.Get(key, nil)
	if err != nil {
		panic(err)
	}
	return string(pair.Value)
}
