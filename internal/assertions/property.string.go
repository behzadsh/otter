package assertions

import "fmt"

type StringFormat uint8

var stringFormatValues = []string{
	"date",
	"date-time",
	"uri",
	"email",
	"ipv4",
	"ipv6",
	"uuid",
}

func (f StringFormat) String() string {
	return stringFormatValues[f]
}

const (
	StringFormatDate StringFormat = iota
	StringFormatDateTime
	StringFormatURI
	StringFormatEmail
	StringFormatIPV4
	StringFormatIPV6
	StringFormatUUID
)

func CastStringFormat(s string) (StringFormat, error) {
	for i, value := range stringFormatValues {
		if s == value {
			return StringFormat(i), nil
		}
	}

	return 0, fmt.Errorf("invalid string format")
}

type StringProperty struct {
	MinLength  *int          `yaml:"min-length"`
	MaxLength  *int          `yaml:"max-length"`
	Length     *int          `yaml:"length"`
	Pattern    *string       `yaml:"pattern"`
	Format     *StringFormat `yaml:"format"`
	DateFormat *string       `yaml:"date-format"`
	Enum       []string      `yaml:"enum"`
}

func (StringProperty) PropertyType() SchemaPropertyType {
	return SchemaPropertyTypeString
}
