package gen

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

var definition = _Definition{
	OpenAPI: "3.0.0",
	Info: _Info{
		Title:   "Example",
		Version: "0.1.0",
	},
	Servers: []_Server{
		{URL: "https://example.com"},
	},
	Paths: map[string]_Path{},
	Components: _Components{
		Schemas: map[string]_Schema{},
	},
	Security: []_Security{},
	Tags:     []_Tag{},
}

func GenerateOpenAPI(config Config, format OpenAPIFormat) string {
	/// Schemas
	schemas := definition.Components.Schemas //map[string]_Schema{}
	for _, v := range config.Models {
		properties := map[string]_Property{}
		for _, attr := range v.Attributes {
			var attrType string
			if strings.Contains(attr.Type, "int") {
				attrType = "integer"
			} else if strings.Contains(attr.Type, "float") {
				attrType = "number"
			} else if attr.Type == "string" || strings.Contains(attr.Type, "ObjectID") || strings.Contains(attr.Type, "Time") {
				attrType = "string"
			} else {
				attrType = "object"
			}

			properties[attr.Name] = _Property{Type: attrType}
		}

		schemas[fmt.Sprintf("%s.%s", config.Collection, v.Name)] = _Schema{
			Type:       "object",
			Properties: properties,
		}
	}

	/// Routes
	paths := definition.Paths
	for _, g := range config.Routes {
		definition.Tags = append(definition.Tags, _Tag{Name: g.Path})
		for _, r := range g.Children {
			re := regexp.MustCompile(`:(\w+)`)
			endpoint := re.ReplaceAllString(fmt.Sprintf("%s%s", g.Path, r.Path), `{$1}`)
			path, ok := paths[endpoint]
			if !ok {
				path = _Path{}
			}

			operation := _Operation{
				Summary: r.Handler,
				Tags:    []string{g.Path},
				Responses: map[int]_Response{
					200: {
						Description: "OK",
						Content: map[string]_MediaType{
							"application/json": {
								Schema: _Schema{
									Type: "object",
								},
							},
						},
					},
				},
			}

			switch r.Method {
			case http.MethodGet:
				path.Get = operation
			case http.MethodPut:
				path.Put = operation
			case http.MethodPost:
				path.Post = operation
			case http.MethodDelete:
				path.Delete = operation
			case http.MethodOptions:
				path.Options = operation
			case http.MethodHead:
				path.Head = operation
			case http.MethodPatch:
				path.Patch = operation
			case http.MethodTrace:
				path.Trace = operation
			}

			paths[endpoint] = path
		}
	}

	if format == JSON {
		data, _ := json.Marshal(definition)
		return string(data)
	} else if format == YML {
		data, _ := yaml.Marshal(definition)
		return string(data)
	} else {
		return "Requires JSON/YML"
	}
}
