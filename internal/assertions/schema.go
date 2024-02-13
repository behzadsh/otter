package assertions

import "fmt"

type SchemaType uint8

var schemaTypeValues = []string{
	"array",
	"object",
}

func (t SchemaType) String() string {
	return schemaTypeValues[t]
}

const (
	SchemaTypeArray SchemaType = iota
	SchemaTypeObject
)

func CastSchemaType(s string) (SchemaType, error) {
	for i, value := range schemaTypeValues {
		if s == value {
			return SchemaType(i), nil
		}
	}

	return 0, fmt.Errorf("invalid schema type")
}

type SchemaAssertion interface {
	GetType() SchemaType
	Validate() error
}
