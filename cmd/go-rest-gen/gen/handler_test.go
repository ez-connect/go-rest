package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateHandler(t *testing.T) {
	v := GenerateHandler(
		"test",
		Config{},
	)

	assert.NotEmpty(t, v)
}
