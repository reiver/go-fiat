package fiat

import (
	"testing"
)

func TestDecoderBoolFromBool(t *testing.T) {

	tests := []struct{
		Value bool
	}{
		{
			Value: false,
		},
		{
			Value: true,
		},
	}


	for testNumber, test := range tests {

		var decoder Decoder

		const targetKey = "target_key"

		decoder.data = map[string][]interface{}{}
		decoder.data[targetKey] = []interface{}{ test.Value }

		x, err := decoder.Bool(targetKey)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := bool(x)

		if expected, actual := test.Value, y; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
	}
}
