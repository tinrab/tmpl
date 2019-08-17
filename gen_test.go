package tmpl

import (
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
