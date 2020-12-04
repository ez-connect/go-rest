package gen

//
// https://swagger.io/specification/
//

type OpenAPIFormat string

const (
	JSON OpenAPIFormat = "json"
	YML  OpenAPIFormat = "yml"
)

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
	TermsOfService string   `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        _Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        _License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string   `json:"version,omitempty" yaml:"version,omitempty"`
}

type _ServerVariable struct {
	Enum        []string `json:"enum" yaml:"enum"`
	Default     string   `json:"default" yaml:"enum"`
	Description string   `json:"description" yaml:"enum"`
}

type _Server struct {
	URL         string                     `json:"url" yaml:"url"`
	Description string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Variables   map[string]_ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
}

type _ExternalDoc struct {
	Description string `json:"description" yaml:"description"`
	Url         string `json:"url" yaml:"url"`
}

type _Parameter struct {
	// Reference Object
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`

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
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`

	Type       string               `json:"type,omitempty" yaml:"type,omitempty"`
	Required   []string             `json:"required,omitempty" yaml:"required,omitempty"`
	Properties map[string]_Property `json:"properties,omitempty" yaml:"properties,omitempty"`
}

type _MediaType struct {
	Schema   _Schema     `json:"schema" yaml:"schema"`
	Example  interface{} `json:"example,omitempty" yaml:"example,omitempty"`
	Examples interface{} `json:"examples,omitempty" yaml:"examples,omitempty"`
	Encoding interface{} `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

type _RequestBody struct {
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`

	Description string     `json:"description" yaml:"description"`
	Content     _MediaType `json:"content" yaml:"content"`
	Required    bool       `json:"required" yaml:"required"`
}

type _Response struct {
	Description string                `json:"description" yaml:"description"`
	Headers     string                `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]_MediaType `json:"content,omitempty" yaml:"content,omitempty"`
}

type _Operation struct {
	Tags         []string          `json:"tags" yaml:"tags"`
	Summary      string            `json:"summary" yaml:"summary"`
	Description  string            `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs _ExternalDoc      `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationId  string            `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   []_Parameter      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Requestbody  _RequestBody      `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    map[int]_Response `json:"responses" yaml:"responses"`
}

type _Path struct {
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`

	Summary     string       `json:"summary" yaml:"summary"`
	Description string       `json:"description,omitempty" yaml:"description,omitempty"`
	Get         _Operation   `json:"get,omitempty" yaml:"get,omitempty"`
	Put         _Operation   `json:"put,omitempty" yaml:"put,omitempty"`
	Post        _Operation   `json:"post,omitempty" yaml:"post,omitempty"`
	Delete      _Operation   `json:"delete,omitempty" yaml:"delete,omitempty"`
	Options     _Operation   `json:"options,omitempty" yaml:"options,omitempty"`
	Head        _Operation   `json:"head,omitempty" yaml:"head,omitempty"`
	Patch       _Operation   `json:"patch,omitempty" yaml:"patch,omitempty"`
	Trace       _Operation   `json:"trace,omitempty" yaml:"trace,omitempty"`
	Servers     []_Server    `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters  []_Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

type _Components struct {
	Schemas map[string]_Schema `json:"schemas" yaml:"schemas"`
}

type _Security struct {
}

type _Tag struct {
	Name string `json:"name" yaml:"name"`
}

type _Definition struct {
	OpenAPI      string           `json:"openapi" yaml:"openapi"`
	Info         _Info            `json:"info" yaml:"info"`
	Servers      []_Server        `json:"servers" yaml:"servers"`
	Paths        map[string]_Path `json:"paths" yaml:"paths"`
	Components   _Components      `json:"components" yaml:"components"`
	Security     []_Security      `json:"security" yaml:"security"`
	Tags         []_Tag           `json:"tags" yaml:"tags"`
	ExternalDocs _ExternalDoc     `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}
