package assertions

type Assertion struct {
	Schema []SchemaAssertion `yaml:"schema"`
	Values []ValueAssertion  `yaml:"values"`
}
