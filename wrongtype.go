package fiat

import (
	"fmt"
)

type WrongTypeComplainer interface {
	error
	WrongTypeComplainer()
}

type internalWrongTypeComplainer struct{
	expectedType string
	actualType string
}

func (receiver internalWrongTypeComplainer) Error() string {
	return fmt.Sprintf("Wrong Type: %q; expected something compatible with: %q", receiver.actualType, receiver.expectedType)
}

func (internalWrongTypeComplainer) WrongTypeComplainer() {
	// Nothing here.
}
