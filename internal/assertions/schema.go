package assertions

import "fmt"

type StructureType uint8

var structureTypeValues = []string{
	"array",
	"object",
}

func (t StructureType) String() string {
	return structureTypeValues[t]
}

const (
	StructureTypeArray StructureType = iota
	StructureTypeObject
)

func CastStructureType(s string) (StructureType, error) {
	for i, value := range structureTypeValues {
		if s == value {
			return StructureType(i), nil
		}
	}

	return 0, fmt.Errorf("invalid structure type")
}

type SchemaAssertion interface {
	GetType() StructureType
	Validate() error
}

type ObjectStructureAssertion struct {
	Properties any `yaml:"properties"`
}
