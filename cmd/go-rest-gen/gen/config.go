package gen

const MainModelName = "Model"

// Import returns all imports of all files
type Import struct {
	Model      []string `yaml:"model,omitempty"`
	Repository []string `yaml:"repository,omitempty"`
	Handler    []string `yaml:"handler,omitempty"`
	Router     []string `yaml:"router,omitempty"`
}

///////////////////////////////////////////////////////////////////

type Attribute struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Validate    string `yaml:"validate,omitempty"`
	AllowsEmpty bool   `yaml:"allowsEmpty,omitempty"`
}

type ModelConfig struct {
	// Default Model for the collection, using name for embed models
	Name string `yaml:"name"`

	// All attributes
	Attributes []Attribute `yaml:"attributes"`
}

///////////////////////////////////////////////////////////////////

// Single index
type SingleIndex struct {
	Field  string `yaml:"field"`
	Order  int    `yaml:"order,omitempty"`
	Unique bool   `yaml:"unique,omitempty"`
}

// Will support compound order
// For a single-field index and sort operations,
// the sort order (i.e. ascending or descending)
// of the index key does not matter because MongoDB
// can traverse the index in either direction.
// https://docs.mongodb.com/manual/indexes/
type CompoundIndexField struct {
	Field string `yaml:"field"`
	Order int    `yaml:"order,omitempty"`
}

type CompoundIndex struct {
	Fields []CompoundIndexField `yaml:"fields,omitempty"`
	Unique bool                 `yaml:"unique,omitempty"`
}

type Index struct {
	Singles   []SingleIndex   `yaml:"singles,omitempty"`
	Compounds []CompoundIndex `yaml:"compounds,omitempty"`
	Texts     []string        `yaml:"texts,omitempty"`
}

///////////////////////////////////////////////////////////////////

type ParameterIn string

const (
	ParameterInPath  ParameterIn = "path"
	ParameterInQuery ParameterIn = "query"
	ParameterInBody  ParameterIn = "body"
)

type RouteParameter struct {
	Name     string
	In       ParameterIn
	Type     string
	Required bool `yaml:"required,omitempty"`
}

type RouteConfig struct {
	Path           string
	Method         string
	Parameters     []RouteParameter
	Handler        string
	MiddlewareFunc string

	// Permission / Policies here
}

type RouteGroup struct {
	Path           string
	MiddlewareFunc string
	Children       []RouteConfig
}

///////////////////////////////////////////////////////////////////

type Config struct {
	Import     Import        `yaml:"import"`
	Collection string        `yaml:"collection"`
	Models     []ModelConfig `yaml:"models"`
	Index      Index         `yaml:"index"`
	Routes     []RouteGroup  `yaml:"routes"`
	LifeCycle  string        `yaml:"lifeCycle"`
}
