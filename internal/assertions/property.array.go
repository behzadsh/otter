package assertions

type ArrayProperty struct {
	Items       SchemaProperty `yaml:"items"`
	Length      *int           `yaml:"length"`
	MinItems    *int           `yaml:"min-items"`
	MaxItems    *int           `yaml:"max-items"`
	UniqueItems bool           `yaml:"unique-items"`
}

func (ArrayProperty) PropertyType() SchemaPropertyType {
	return SchemaPropertyTypeArray
}
