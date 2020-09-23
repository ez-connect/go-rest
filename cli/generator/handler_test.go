package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateHandler(t *testing.T) {
	v := GenerateHandler(
		"test",
	)

	assert.NotEmpty(t, v)
}
