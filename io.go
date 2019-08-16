package tmpl

import (
	"os"
	"path/filepath"
)

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
