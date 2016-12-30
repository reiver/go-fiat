package fiatdriver

import (
	"fmt"
)

type NotFoundComplainer interface {
	error
	NotFoundComplainer()
}

type internalNotFoundComplainer struct{
	name string
}

func (receiver internalNotFoundComplainer) Error() string {
	return fmt.Sprintf("Not Found: %q", receiver.name)
}

func (internalNotFoundComplainer) NotFoundComplainer() {
	// Nothing here.
}
