package generator

import "encoding/json"

type _ReferenceObject struct {
	Ref string `json:"$ref,omitempty"`
}

type _Contact struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Email string `json:"email"`
}

type _License struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type _Info struct {
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	TermsOfService string   `json:"termsOfService"`
	Contact        _Contact `json:"contact"`
	License        _License `json:"license"`
	Version        string   `json:"version"`
}

type _ServerVariable struct {
	Enum        []string `json:"enum"`
	Default     string   `json:"default"`
	Description string   `json:"description"`
}

type _Server struct {
	URL         string                     `json:"url"`
	Description string                     `json:"description"`
	Variables   map[string]_ServerVariable `json:"variables"`
}

type _ExternalDoc struct {
	Description string `json:"description"`
	Url         string `json:"url"`
}

type _Parameter struct {
	// Reference Object
	_ReferenceObject

	// Or, Parameter Object
	Name            string `json:"name"`
	In              string `json:"in"`
	Description     string `json:"description"`
	Required        bool   `json:"required,omitempty"`
	Deprecated      bool   `json:"deprecated,omitempty"`
	AllowEmptyValue bool   `json:"allowEmptyValue,omitempty"`
}

type _Property struct {
	Type string `json:"type"`
}

type _Schema struct {
	_ReferenceObject

	Type       string               `json:"type"`
	Required   []string             `json:"required"`
	Properties map[string]_Property `json:"properties"`
}

type _MediaType struct {
	Schema   _Schema     `json:"schema"`
	Example  interface{} `json:"example"`
	Examples interface{} `json:"examples"`
	Encoding interface{} `json:"encoding"`
}

type _RequestBody struct {
	_ReferenceObject

	Description string     `json:"description"`
	Content     _MediaType `json:"content"`
	Required    bool       `json:"required"`
}

type _Operation struct {
	Tags         []string     `json:"tags"`
	Summary      string       `json:"summary"`
	Description  string       `json:"description"`
	ExternalDocs _ExternalDoc `json:"externalDocs"`
	OperationId  string       `json:"operationId"`
	Parameters   []_Parameter `json:"parameters"`
	Requestbody  _RequestBody `json:"requestBody"`
}

type _Path struct {
	_ReferenceObject

	Summary     string     `json:"summary"`
	Description string     `json:"description"`
	Get         _Operation `json:"get"`
}

type _Component struct {
	Schemas []_Schema `json:"schemas"`
}

type _Security struct {
}

type _Tag struct {
}

type _Definition struct {
	OpenAPI      string           `json:"openapi"`
	Info         _Info            `json:"info"`
	Servers      []_Server        `json:"servers"`
	Paths        map[string]_Path `json:"paths"`
	Components   _Component       `json:"components"`
	Security     []_Security      `json:"security"`
	Tags         []_Tag           `json:"tags"`
	ExternalDocs []_ExternalDoc   `json:"externalDocs"`
}

func GenerateOpenAPI(packageName string, collections []string) string {
	buf := _Definition{
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

	data, _ := json.Marshal(buf)
	return string(data)
}
