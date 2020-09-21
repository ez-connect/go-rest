package generator

type Attribute struct {
	Type     string
	Required bool
}

type ModelConfig struct {
	// Collection name
	Name string
	// All attributes
	Attributes map[string]Attribute
}
