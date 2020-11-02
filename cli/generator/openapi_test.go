package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOpenAPI(t *testing.T) {
	v := GenerateOpenAPIJSON("", []string{})
	t.Error(v)
	assert.NotEmpty(t, v)

	v = GenerateOpenAPIYML("", []string{})
	t.Error(v)
	assert.NotEmpty(t, v)
}
