package gen

import (
	"log"
	"os"
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
	}, _YML)

	f, err := os.Create("openapi_test.yml")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(v)
	if err != nil {
		log.Fatal(err)
	}

	assert.NotEmpty(t, v)
}
