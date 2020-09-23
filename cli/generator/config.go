package generator

type Attribute struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Required bool   `yaml:"required"`
}

type ModelConfig struct {
	// Collection name
	Name string `yaml:"name"`

	// All attributes
	Attributes []Attribute `yaml:"attributes"`
}

type RouteGroup struct {
	Path     string
	Children []RouteConfig
}

type RouteConfig struct {
	Method  string
	Path    string
	Handler string

	// Permission / Policies here
}

type Config struct {
	Model  ModelConfig
	Routes []RouteGroup
}
