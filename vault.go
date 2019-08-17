package tmpl

import (
	"errors"
	"strings"
	"text/template"

	"github.com/hashicorp/vault/api"
)

type VaultConfig struct {
	Address string
	Token   string
}

func buildVaultFunctionMap(options Options, m template.FuncMap) error {
	if options.VaultConfig == nil {
		return nil
	}
	opts := options.VaultConfig
	cfg := api.DefaultConfig()
	cfg.Address = opts.Address
	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	if opts.Token != "" {
		client.SetToken(opts.Token)
	}

	m["vault"] = func(path, key string) (interface{}, error) {
		secret, err := client.Logical().Read(path)
		if err != nil {
			return nil, err
		}
		if secret == nil {
			return nil, nil
		}
		if len(secret.Warnings) != 0 {
			return nil, errors.New(secret.Warnings[0])
		}
		data, ok := secret.Data["data"].(map[string]interface{})
		if !ok {
			return nil, nil
		}
		return getMapValue(data, key), nil
	}

	return nil
}

func getMapValue(m map[string]interface{}, path string) interface{} {
	steps := strings.Split(path, "/")
	var value interface{} = m
	for _, step := range steps {
		vm, ok := value.(map[string]interface{})
		if !ok {
			break
		}
		if step == "" {
			return vm
		}
		if v, ok := vm[step]; ok {
			value = v
		} else {
			return nil
		}
	}
	return value
}
