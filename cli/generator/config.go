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

type Index struct {
	Fields []string
	Order  int // 1: asc or -1: desc
	Unique bool
	Text   bool
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
	Model   ModelConfig
	Indexes []Index
	Routes  []RouteGroup
}
