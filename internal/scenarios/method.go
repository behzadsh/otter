package scenarios

import "fmt"

type Method uint8

var methodValues = []string{
	"GET",
	"POST",
	"PUT",
	"PATCH",
	"DELETE",
}

func (m Method) String() string {
	return methodValues[m]
}

const (
	MethodGet Method = iota
	MethodPost
	MethodPut
	MethodPatch
	MethodDelete
)

func CastMethod(s string) (Method, error) {
	for i, value := range methodValues {
		if s == value {
			return Method(i), nil
		}
	}

	return 0, fmt.Errorf("invalid method")
}
