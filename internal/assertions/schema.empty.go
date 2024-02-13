package assertions

type EmptySchemaAssertion struct{}

func (e EmptySchemaAssertion) GetType() SchemaType {
	return SchemaTypeEmpty
}

func (e EmptySchemaAssertion) Validate() error {
	// TODO implement me
	panic("implement me")
}
