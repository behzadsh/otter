package assertions

type ObjectSchemaAssertion struct {
	Properties map[string]SchemaProperty `yaml:"properties"`
	Required   []string                  `yaml:"required"`
}

func (ObjectSchemaAssertion) GetType() SchemaType {
	return SchemaTypeObject
}

func (ObjectSchemaAssertion) Validate() error {
	// TODO implement me
	panic("implement me")
}
