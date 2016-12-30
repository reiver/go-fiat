package fiat

import (
	"errors"
)

var (
	errInternalError    = errors.New("Internal Error")
	errNilReflectedType = errors.New("Nil Reflected Type")
	errNilReceiver      = errors.New("Nil Receiver")
	errRuneError        = errors.New("Rune Error")
)
