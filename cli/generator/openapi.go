package generator

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type _ReferenceObject struct {
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

type _Contact struct {
	Name  string `json:"name" yaml:"name"`
	URL   string `json:"url" yaml:"url"`
	Email string `json:"email" yaml:"email"`
}

type _License struct {
	Name string `json:"name" yaml:"name"`
	URL  string `json:"url" yaml:"url"`
}

type _Info struct {
	Title          string   `json:"title" yaml:"title"`
	Description    string   `json:"description" yaml:"description"`
	TermsOfService string   `json:"termsOfService" yaml:"termsOfService"`
	Contact        _Contact `json:"contact" yaml:"contact"`
	License        _License `json:"license" yaml:"license"`
	Version        string   `json:"version" yaml:"version"`
}

type _ServerVariable struct {
	Enum        []string `json:"enum" yaml:"enum"`
	Default     string   `json:"default" yaml:"enum"`
	Description string   `json:"description" yaml:"enum"`
}

type _Server struct {
	URL         string                     `json:"url" yaml:"url"`
	Description string                     `json:"description" yaml:"description"`
	Variables   map[string]_ServerVariable `json:"variables" yaml:"variables"`
}

type _ExternalDoc struct {
	Description string `json:"description" yaml:"description"`
	Url         string `json:"url" yaml:"url"`
}

type _Parameter struct {
	// Reference Object
	_ReferenceObject

	// Or, Parameter Object
	Name            string `json:"name" yaml:"name"`
	In              string `json:"in" yaml:"in"`
	Description     string `json:"description" yaml:"description"`
	Required        bool   `json:"required,omitempty" yaml:"required"`
	Deprecated      bool   `json:"deprecated,omitempty" yaml:"deprecated"`
	AllowEmptyValue bool   `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue"`
}

type _Property struct {
	Type string `json:"type"`
}

type _Schema struct {
	_ReferenceObject

	Type       string               `json:"type" yaml:"type"`
	Required   []string             `json:"required" yaml:"required"`
	Properties map[string]_Property `json:"properties" yaml:"properties"`
}

type _MediaType struct {
	Schema   _Schema     `json:"schema" yaml:"schema"`
	Example  interface{} `json:"example" yaml:"example"`
	Examples interface{} `json:"examples" yaml:"examples"`
	Encoding interface{} `json:"encoding" yaml:"encoding"`
}

type _RequestBody struct {
	_ReferenceObject

	Description string     `json:"description" yaml:"description"`
	Content     _MediaType `json:"content" yaml:"content"`
	Required    bool       `json:"required" yaml:"required"`
}

type _Operation struct {
	Tags         []string     `json:"tags" yaml:"tags"`
	Summary      string       `json:"summary" yaml:"summary"`
	Description  string       `json:"description" yaml:"description"`
	ExternalDocs _ExternalDoc `json:"externalDocs" yaml:"externalDocs"`
	OperationId  string       `json:"operationId" yaml:"operationId"`
	Parameters   []_Parameter `json:"parameters" yaml:"parameters"`
	Requestbody  _RequestBody `json:"requestBody" yaml:"requestBody"`
}

type _Path struct {
	_ReferenceObject

	Summary     string     `json:"summary" yaml:"summary"`
	Description string     `json:"description" yaml:"description"`
	Get         _Operation `json:"get" yaml:"get"`
}

type _Component struct {
	Schemas []_Schema `json:"schemas" yaml:"schemas"`
}

type _Security struct {
}

type _Tag struct {
}

type _Definition struct {
	OpenAPI      string           `json:"openapi" yaml:"openapi"`
	Info         _Info            `json:"info" yaml:"info"`
	Servers      []_Server        `json:"servers" yaml:"servers"`
	Paths        map[string]_Path `json:"paths" yaml:"paths"`
	Components   _Component       `json:"components" yaml:"components"`
	Security     []_Security      `json:"security" yaml:"security"`
	Tags         []_Tag           `json:"tags" yaml:"tags"`
	ExternalDocs []_ExternalDoc   `json:"externalDocs" yaml:"externalDocs"`
}

var kDefaultAPI = _Definition{
	OpenAPI: "3.0.0",
	Info: _Info{
		Title: "a",
	},
	Servers: []_Server{
		{URL: "https://example.com"},
	},
	Paths:        map[string]_Path{},
	Security:     []_Security{},
	ExternalDocs: []_ExternalDoc{},
}

func GenerateOpenAPIJSON(packageName string, collections []string) string {
	data, _ := json.Marshal(kDefaultAPI)
	return string(data)
}

func GenerateOpenAPIYML(packageName string, collections []string) string {
	data, _ := yaml.Marshal(kDefaultAPI)
	return string(data)
}
