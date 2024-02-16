package assertions

type NumberProperty struct {
	Min        *float64 `yaml:"min"`
	Max        *float64 `yaml:"max"`
	MultipleOf *float64 `yaml:"multiple-of"`
	Equals     *float64 `yaml:"equals"`
}

func (NumberProperty) PropertyType() SchemaPropertyType {
	return SchemaPropertyTypeNumber
}
