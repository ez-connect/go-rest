package generator

type Attribute struct {
	Name     string
	Type     string
	Required bool
}

type ModelConfig struct {
	// Collection name
	Name string
	// All attributes
	Attributes []Attribute
}
