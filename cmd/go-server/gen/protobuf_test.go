package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateProtobuf(t *testing.T) {
	v := GenerateProtobuf(
		"test",
		Config{
			Collection: "test",
			Models: []ModelConfig{
				{
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
			},
		},
	)

	t.Error(v)
	assert.NotEmpty(t, v)
}
