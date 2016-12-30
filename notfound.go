package fiat

import (
	"fmt"
)

type NotFoundComplainer interface {
	error
	NotFoundComplainer()
}

type internalNotFoundComplainer struct{
	name string
	have []string
}

func (receiver internalNotFoundComplainer) Error() string {
	return fmt.Sprintf("Not Found: %q; %v", receiver.name, receiver.have)
}

func (internalNotFoundComplainer) NotFoundComplainer() {
	// Nothing here.
}
