package assertions

type ObjectProperty struct {
	Properties map[string]SchemaProperty `yaml:"properties"`
	Required   []string                  `yaml:"required"`
}

func (ObjectProperty) PropertyType() SchemaPropertyType {
	return SchemaPropertyTypeObject
}
