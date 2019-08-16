package tmpl

import (
	"crypto/md5"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	opts := Options{
		Parameters: map[string]interface{}{
			"name": "John",
			"age":  14000,
		},
		Sources: []Source{{
			Name:  "a.txt",
			Value: "{{.name}}, {{.age}}",
		}},
	}

	results, err := Generate(opts)
	assert.NoError(t, err)

	assert.Equal(t, "a.txt", results[0].Name)
	assert.Equal(t, "John, 14000", string(results[0].Data))
}

func TestGenerateFromFiles(t *testing.T) {
	results, err := GenerateFromFiles("./test/*.params.*", "./test/*.tmpl.yaml")
	assert.NoError(t, err)

	cases := []struct {
		name  string
		value string
	}{{
		name:  "test/deployment.tmpl.yaml",
		value: "I4g2JuJyFEd/Z8AJeEdptw",
	}, {
		name:  "test/service.tmpl.yaml",
		value: "/cm8eQ8bE1TV13Nh3Xh/bg",
	}}

	for i, c := range cases {
		assert.Equal(t, c.name, results[i].Name)
		assert.Equal(t, c.value, hash(results[i].Data))
	}
}

func hash(data []byte) string {
	h := md5.Sum(data)
	return base64.RawStdEncoding.EncodeToString(h[:])
}
