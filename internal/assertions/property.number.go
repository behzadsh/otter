package assertions

type NumberProperty struct {
	Min        *float64 `yaml:"min"`
	Max        *float64 `yaml:"max"`
	MultipleOf *float64 `yaml:"multiple-of"`
}

func (NumberProperty) PropertyType() SchemaPropertyType {
	return SchemaPropertyTypeNumber
}
