package generator

import (
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/yaml.v3"
)

func GenerateSettings(packageName string) string {
	doc := Config{
		Model: ModelConfig{
			Name: packageName,
			Attributes: []Attribute{
				{Name: "name", Type: "string"},
				{Name: "price", Type: "float32"},
			},
		},
		EmbedModels: []ModelConfig{},
		Indexes:     []Index{},
		Routes: []RouteGroup{
			{
				Path: fmt.Sprintf("/%ss", packageName),
				Children: []RouteConfig{
					{
						Method:  http.MethodGet,
						Handler: fmt.Sprintf("Find%s", strings.Title(packageName)),
					},
					{
						Method:  http.MethodPost,
						Handler: fmt.Sprintf("Insert%s", strings.Title(packageName)),
					},
					{
						Method:  http.MethodGet,
						Path:    "/:id",
						Handler: fmt.Sprintf("FindOne%s", strings.Title(packageName)),
					},
					{
						Method:  http.MethodPut,
						Path:    "/:id",
						Handler: fmt.Sprintf("Update%s", strings.Title(packageName)),
					},
					{
						Method:  http.MethodDelete,
						Path:    "/:id",
						Handler: fmt.Sprintf("Delete%s", strings.Title(packageName)),
					},
				},
			},
		},
	}

	data, err := yaml.Marshal(doc)
	if err != nil {
		return ""
	}

	return string(data)
}
