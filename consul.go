package tmpl

import (
	"errors"
	"text/template"

	"github.com/hashicorp/consul/api"
)

type ConsulConfig struct {
	Address string
	Token   string
}

func buildConsulFunctionMap(options Options, m template.FuncMap) error {
	if options.ConsulConfig == nil {
		return nil
	}

	cfg := api.DefaultConfig()
	cfg.Address = options.ConsulConfig.Address
	if options.ConsulConfig.Token != "" {
		cfg.Token = options.ConsulConfig.Token
	}

	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	_, err = client.Status().Leader()
	if err != nil {
		return errors.New("Consul connection error " + cfg.Address)
	}

	kv := client.KV()

	m["consul"] = func(key string) (string, error) {
		pair, _, err := kv.Get(key, &api.QueryOptions{
			UseCache: true,
		})
		if err != nil {
			return "", err
		}
		return string(pair.Value), nil
	}

	return nil
}
