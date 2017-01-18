package fiat

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	randomness = rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))
)


type scannableString struct {
	scanned bool
	s       string
}

func (receiver *scannableString) Scan(v interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("Cannot scan %T into %T", v, *receiver)
	}

	receiver.s       = s
	receiver.scanned = true

	return nil
}

func (receiver scannableString) String() (string, error) {

	if ! receiver.scanned {
		return "", errors.New("Not scanned yet.")
	}

	return receiver.s, nil
}
