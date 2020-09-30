package generator

type Attribute struct {
	Name      string
	Type      string
	Required  bool
	Omitempty bool
}

type ModelConfig struct {
	// Collection name for main model or name for embed model
	Name string

	// All attributes
	Attributes []Attribute
}

type Index struct {
	Fields []string
	Order  int  // 1: asc, -1: desc
	Text   bool // text index
	Unique bool
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
	// embed structures only
	EmbedModels []ModelConfig
}
