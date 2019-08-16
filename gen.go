package tmpl

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"path/filepath"
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
	Parameters ParameterMap
	Sources    []Source
}

func Generate(options Options) ([]Result, error) {
	var results []Result
	for _, src := range options.Sources {
		tmpl, err := template.New("").Parse(src.Value)
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

func GenerateFromFiles(parametersGlob string, sourcesGlob string) ([]Result, error) {
	pFiles, err := filepath.Glob(parametersGlob)
	if err != nil {
		return nil, err
	}
	pm, err := parseParameterMap(pFiles)
	if err != nil {
		return nil, err
	}
	sources, err := readSourceGlob(sourcesGlob)
	if err != nil {
		return nil, err
	}
	options := Options{
		Parameters: pm,
		Sources:    sources,
	}
	return Generate(options)
}

func readSourceGlob(glob string) ([]Source, error) {
	files, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}
	var sources []Source
	for _, f := range files {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}
		sources = append(sources, Source{
			Name:  f,
			Value: string(data),
		})
	}
	return sources, nil
}
