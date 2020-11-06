package gen

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOpenAP(t *testing.T) {
	v := GenerateOpenAPI(Config{
		Collection: "test",
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
						Name: "quantity",
						Type: "int32",
					},
					{
						Name: "price",
						Type: "float32",
					},
					{
						Name: "createdAt",
						Type: "*time.Time",
					},
				},
			},
		},
		Routes: []RouteGroup{
			{
				Path: "/example",
				Children: []RouteConfig{
					{
						Path:    "",
						Method:  "GET",
						Handler: "FindAll",
					},
					{
						Path:    "",
						Method:  "POST",
						Handler: "Insert",
					},
					{
						Path:    "/:id",
						Method:  "GET",
						Handler: "FindOne",
					},
					{
						Path:    "/:id",
						Method:  "PUT",
						Handler: "Update",
					},
					{
						Path:    "/:id",
						Method:  "DELETE",
						Handler: "Delete",
					},
				},
			},
		},
	}, YML)

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
