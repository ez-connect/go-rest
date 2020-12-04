package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRepository(t *testing.T) {
	v := GenerateRepository(
		"test",
		Config{
			Index: Index{
				Singles: []SingleIndex{
					{Field: "name", Unique: true},
				},
			},
		},
	)

	assert.NotEmpty(t, v)
}
