package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOpenAP(t *testing.T) {
	v := GenerateOpenAPI(Config{
		Models: []ModelConfig{
			{Name: "Model",
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
	}, _JSON)

	t.Error(v)
	assert.NotEmpty(t, v)
}
