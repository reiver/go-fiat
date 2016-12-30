package fiatdriver

import (
	"errors"
)

var (
	errFound        = errors.New("Found")
	errNilDriver    = errors.New("Nil Driver")
	errNilReceiver  = errors.New("Nil Receiver")
	errNotFound     = errors.New("Not Found")
)
