package assertions

type Assertion struct {
	Status int               `yaml:"status"`
	Schema []SchemaAssertion `yaml:"schema"`
	Values []ValueAssertion  `yaml:"values"`
}
