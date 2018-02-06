package api

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"os"
)

func (kvStruct *KeyValue) WriteKVsToConsul() error {
	for key, value := range kvStruct.kv {
		if os.Getenv("CONSUL_IP") == "" {
			return errors.New("CONSUL_IP environment variable not set.")
		}
		err := requestPUT(os.Getenv("CONSUL_IP"), key, value)
		if err != nil {
			return err
		}
		fmt.Println("key:", key, "value", value)
	}
	fmt.Println("Wrote KVs to Consul")
	return nil
}

func GetKVFromConsul(key string) (string, error) {
	if os.Getenv("CONSUL_IP") == "" {
		return "", errors.New("CONSUL_IP environment variable not set.")
	}
	resp, err := requestGET(os.Getenv("CONSUL_IP"), key)
	return resp, err
}

func GetKVsFromConsul() ([]string, error) {
	if os.Getenv("CONSUL_IP") == "" {
		return []string{""}, errors.New("CONSUL_IP environment variable not set.")
	}
	resp, err := requestGETS(os.Getenv("CONSUL_IP"))
	return resp, err
}

func requestPUT(url string, key string, value string) error {
	config := api.DefaultConfig()
	config.Address = url + ":8500"
	client, err := api.NewClient(config)

	if err != nil {
		return err
	}

	kv := client.KV()

	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err = kv.Put(p, nil)
	if err != nil {
		return err
	}

	return nil
}

func requestGET(url string, key string) (string, error) {
	config := api.DefaultConfig()
	config.Address = url + ":8500"
	client, err := api.NewClient(config)

	kv := client.KV()

	pair, _, err := kv.Get(key, nil)

	if pair == nil {
		return string("No value found for key."), err
	}
	return string(pair.Value), err

}

func requestGETS(url string) ([]string, error) {
	config := api.DefaultConfig()
	config.Address = url + ":8500"
	client, err := api.NewClient(config)

	kv := client.KV()

	pairs, _, err := kv.List("", nil)

	if len(pairs) == 0 {
		return []string{"No keys found."}, err
	}

	var res []string

	for _, keypair := range pairs {
		res = append(res, keypair.Key)
	}

	return res, err
}
