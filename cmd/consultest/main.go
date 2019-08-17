package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

func main() {
	cfg := api.DefaultConfig()
	client, err := api.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	kv := client.KV()
	appVersion, _, err := kv.Get("app/version", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(appVersion.Value))
}
