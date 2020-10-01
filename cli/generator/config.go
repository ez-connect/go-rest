package generator

type Attribute struct {
	Name      string `yaml:"name"`
	Type      string `yaml:"type"`
	Required  bool   `yaml:"required,omitempty"`
	Omitempty bool   `yaml:"omitempty,omitempty"`
}

type ModelConfig struct {
	// Collection name for main model or name for embed model
	Name string `yaml:"name"`

	// All attributes
	Attributes []Attribute `yaml:"attributes"`
}

type Index struct {
	Fields []string `yaml:"fields"`
	Order  int      `yaml:"order,omitempty"` // 1: asc, -1: desc
	Text   bool     `yaml:"text,omitempty"`  // text index
	Unique bool     `yaml:"unique,omitempty"`
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
	Model       ModelConfig
	EmbedModels []ModelConfig `yaml:"embedModels"`
	Indexes     []Index
	Routes      []RouteGroup
}
