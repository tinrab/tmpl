package tmpl

import (
	"bytes"
	"encoding/json"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Source struct {
	Name  string
	Value string
}

type Result struct {
	Name string
	Data []byte
}

type Options struct {
	Parameters   ParameterMap
	Sources      []Source
	ConsulConfig *ConsulConfig
}

func Generate(options Options) ([]Result, error) {
	fm, err := buildFunctions(options)
	if err != nil {
		return nil, err
	}

	var results []Result
	for _, src := range options.Sources {
		tmpl, err := template.New("").
			Funcs(fm).
			Parse(src.Value)
		if err != nil {
			return nil, err
		}
		out := bytes.Buffer{}
		if err := tmpl.Execute(&out, options.Parameters); err != nil {
			return nil, err
		}
		results = append(results, Result{
			Name: src.Name,
			Data: out.Bytes(),
		})
	}
	return results, nil
}

func buildFunctions(options Options) (template.FuncMap, error) {
	fm := template.FuncMap{}

	if err := buildConsulFunctionMap(options, fm); err != nil {
		return nil, err
	}

	fm["json"] = func(v interface{}) string {
		js, _ := json.Marshal(v)
		return string(js)
	}
	fm["jsonParse"] = func(s string) (map[string]interface{}, error) {
		data := map[string]interface{}{}
		if err := json.Unmarshal([]byte(s), &data); err != nil {
			return nil, err
		}
		return data, nil
	}
	fm["yaml"] = func(v interface{}) string {
		ym, _ := yaml.Marshal(v)
		return string(ym)
	}
	fm["yamlParse"] = func(s string) (map[string]interface{}, error) {
		data := map[string]interface{}{}
		if err := yaml.Unmarshal([]byte(s), &data); err != nil {
			return nil, err
		}
		return data, nil
	}
	return fm, nil
}
