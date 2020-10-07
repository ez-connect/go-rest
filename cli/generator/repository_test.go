package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRepository(t *testing.T) {
	v := GenerateRepository(
		"test",
		Config{
			Indexes: []Index{},
		},
	)

	assert.NotEmpty(t, v)
}
