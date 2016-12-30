package fiatdriver

import (
	"testing"
)

func TestRegistrarInternalRegistrar(t *testing.T) {

	var x Registrar = new(internalRegistrar)

	if nil == x {
		t.Errorf("This should never happen.")
	}
}
