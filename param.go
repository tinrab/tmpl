package tmpl

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"
)

type ParameterMap map[string]interface{}

func parseParameterMap(filepaths []string) (ParameterMap, error) {
	result := ParameterMap{}
	for _, fp := range filepaths {
		data, err := ioutil.ReadFile(fp)
		if err != nil {
			return nil, err
		}
		ext := filepath.Ext(fp)
		pm := ParameterMap{}
		if ext == ".yaml" || ext == ".yml" {
			if err := yaml.Unmarshal(data, &pm); err != nil {
				return nil, err
			}
		} else if ext == ".json" {
			if err := json.Unmarshal(data, &pm); err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("unsupported file type: " + ext)
		}
		if err := mergo.Merge(&result, pm); err != nil {
			return nil, err
		}
	}
	return result, nil
}
