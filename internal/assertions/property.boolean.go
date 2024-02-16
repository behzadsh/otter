package assertions

type BooleanProperty struct {
	Equals bool `yaml:"equals"`
}

func (BooleanProperty) PropertyType() SchemaPropertyType {
	return SchemaPropertyTypeBoolean
}
