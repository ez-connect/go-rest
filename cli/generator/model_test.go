package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateModel(t *testing.T) {
	v := GenerateModel(
		"test",
		ModelConfig{
			Name: "ImAModel",
			Attributes: []Attribute{
				{
					Name: "id",
					Type: "*primitive.ObjectID",
				},
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "createdAt",
					Type: "*time.Time",
				},
			},
		},
		[]ModelConfig{},
	)

	assert.NotEmpty(t, v)
}
