package tmpl

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func LoadParameters(path string) (ParameterMap, error) {
	parameterFiles, err := findFiles(path)
	if err != nil {
		return nil, err
	}
	pm, err := parseParameterMap(parameterFiles)
	if err != nil {
		return nil, err
	}
	return pm, err
}

func LoadSources(path string) ([]Source, error) {
	sourceFiles, err := findFiles(path)
	if err != nil {
		return nil, err
	}
	var sources []Source
	for _, f := range sourceFiles {
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

func findFiles(path string) ([]string, error) {
	files, err := filepath.Glob(path)
	if err != nil {
		return nil, err
	}
	n := 0
	for _, f := range files {
		fi, err := os.Stat(f)
		if err != nil {
			return nil, err
		}
		if !fi.IsDir() {
			files[n] = f
			n++
		}
	}
	files = files[:n]
	return files, nil
}
