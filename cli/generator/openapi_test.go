package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOpenAPI(t *testing.T) {
	v := GenerateOpenAPI("", []string{})
	t.Error(v)
	assert.NotEmpty(t, v)
}
