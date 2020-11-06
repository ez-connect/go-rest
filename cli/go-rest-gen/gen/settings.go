package gen

import (
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/yaml.v3"
)

func GenerateSettings(packageName string) string {
	doc := Config{
		Collection: fmt.Sprintf("%ss", packageName),
		Models: []ModelConfig{
			{
				Name: MainModelName,
				Attributes: []Attribute{
					{Name: "name", Type: "string"},
					{Name: "price", Type: "float32"},
				},
			},
		},
		Index: Index{
			Singles: []SingleIndex{
				{Field: "name", Unique: true},
				{Field: "price", Order: -1},
			},
			Compounds: []CompoundIndex{
				{
					Fields: []CompoundIndexField{
						{Field: "name", Order: 1},
						{Field: "price", Order: -1},
					},
				},
			},
			Texts: []string{"name", "price"},
		},
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
