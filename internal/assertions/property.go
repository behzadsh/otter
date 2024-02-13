package assertions

import "fmt"

type SchemaPropertyType uint8

var schemaPropertyTypeValues = []string{
	"string",
	"number",
	"object",
	"array",
	"boolean",
	"null",
}

func (t SchemaPropertyType) String() string {
	return schemaPropertyTypeValues[t]
}

const (
	SchemaPropertyTypeString SchemaPropertyType = iota
	SchemaPropertyTypeNumber
	SchemaPropertyTypeObject
	SchemaPropertyTypeArray
	SchemaPropertyTypeBoolean
	SchemaPropertyTypeNull
)

func CastSchemaPropertyType(s string) (SchemaPropertyType, error) {
	for i, value := range schemaPropertyTypeValues {
		if s == value {
			return SchemaPropertyType(i), nil
		}
	}

	return 0, fmt.Errorf("invalid schema property type")
}

type SchemaProperty interface {
	PropertyType() SchemaPropertyType
}
