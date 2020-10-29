package generator

// Import returns all imports of all files
type Import struct {
	Model      []string `yaml:"model,omitempty"`
	Repository []string `yaml:"repository,omitempty"`
	Handler    []string `yaml:"handler,omitempty"`
	Router     []string `yaml:"router,omitempty"`
}

type Attribute struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Required bool   `yaml:"required,omitempty"`
	// Omitempty bool   `yaml:"omitempty,omitempty"`
}

type ModelConfig struct {
	// Collection name for main model or name for embed model
	Name string `yaml:"name"`

	// All attributes
	Attributes []Attribute `yaml:"attributes"`
}

// Will support compound order
// For a single-field index and sort operations,
// the sort order (i.e. ascending or descending)
// of the index key does not matter because MongoDB
// can traverse the index in either direction.
// https://docs.mongodb.com/manual/indexes/
type CompoundIndex struct {
	Field string
	Order int
}

type Index struct {
	Name   string   `yaml:"name,omitempty"`
	Fields []string `yaml:"fields"`
	Text   bool     `yaml:"text,omitempty"` // text index
	Unique bool     `yaml:"unique,omitempty"`
}

type RouteGroup struct {
	Path           string
	MiddlewareFunc string
	Children       []RouteConfig
}

type RouteConfig struct {
	Method  string
	Path    string
	Handler string

	// Permission / Policies here
}

type Config struct {
	Import      Import        `yaml:"import"`
	Model       ModelConfig   `yaml:"model"`
	EmbedModels []ModelConfig `yaml:"embedModels"`
	Indexes     []Index       `yaml:"indexes"`
	Routes      []RouteGroup  `yaml:"routes"`
	LifeCycle   string        `yaml:"lifeCycle"`
}
