package assertions

type ArraySchemaAssertion struct {
	Items       SchemaProperty `yaml:"items"`
	Length      *int           `yaml:"length"`
	MinItems    *int           `yaml:"min-items"`
	MaxItems    *int           `yaml:"max-items"`
	UniqueItems bool           `yaml:"unique-items"`
}

func (ArraySchemaAssertion) GetType() SchemaType {
	return SchemaTypeArray
}

func (ArraySchemaAssertion) Validate() error {
	// TODO implement me
	panic("implement me")
}
