package assertions

type BooleanProperty struct{}

func (BooleanProperty) PropertyType() SchemaPropertyType {
	return SchemaPropertyTypeBoolean
}
