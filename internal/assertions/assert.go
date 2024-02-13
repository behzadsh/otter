package assertions

type Assertion struct {
	Schema []SchemaAssertion `yaml:"structure"`
	Values []ValueAssertion  `yaml:"values"`
}
