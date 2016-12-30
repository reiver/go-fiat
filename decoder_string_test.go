package fiat

import (
	"fmt"
	"unicode/utf8"

	"testing"
)

func TestDecoderStringFromBytes(t *testing.T) {

	tests := []struct{
		Value []byte
	}{
		{
			Value: []byte(""),
		},



		{
			Value: []byte("apple"),
		},
		{
			Value: []byte("banana"),
		},
		{
			Value: []byte("cherry"),
		},



		{
			Value: []byte("apple banana cherry"),
		},



		{
			Value: []byte("😏 💀👻👽 😊 😍 🙂🙁"),
		},
	}

	{
		const numRand = 20
		for i:=0; i<numRand; i++ {
			test := struct{
				Value []byte
			}{
				Value: []byte(fmt.Sprintf("%d", randomness.Int63())),
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		var decoder Decoder

		const targetKey = "target_key"

		decoder.data = map[string][]interface{}{}
		decoder.data[targetKey] = []interface{}{ test.Value }

		x, err := decoder.String(targetKey)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := []byte(x)

		if expected, actual := len(test.Value), len(y); expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
		for index:=0; index<len(y);index++ {
			if expected, actual := test.Value[index], y[index]; expected != actual {
				t.Errorf("For test #%d and index %d, expected %v, but actually got %v.", testNumber, index, expected, actual)
				continue
			}
		}
	}
}

func TestDecoderStringFromRune(t *testing.T) {

	tests := []struct{
		Value rune
	}{
		{
			Value: utf8.MaxRune,
		},
		{
			Value: 56270, // This is an invaid rune.
		},



		{
			Value: '😏',
		},
		{
			Value: '💀',
		},
		{
			Value: '👻',
		},
		{
			Value: '👽',
		},
		{
			Value: '😊',
		},
		{
			Value: '😍',
		},
		{
			Value: '🙂',
		},
		{
			Value: '🙁',
		},
	}
	for r:=rune(0); r<rune(256); r++ {
		test := struct{
			Value rune
		}{
			Value: r,
		}
		tests = append(tests, test)
	}

	{
		const numRand = 20
		for i:=0; i<numRand; i++ {
			test := struct{
				Value rune
			}{
				Value: rune(randomness.Int63n(utf8.MaxRune)),
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		// Skip invalid runes.
		if !utf8.ValidRune(test.Value) {
			continue
		}

		var decoder Decoder

		const targetKey = "target_key"

		decoder.data = map[string][]interface{}{}
		decoder.data[targetKey] = []interface{}{ test.Value }

		x, err := decoder.String(targetKey)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y, n := utf8.DecodeRuneInString(x)
		if expected, actual := len(x), n; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}

		if expected, actual := test.Value, y; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
	}
}

func TestDecoderStringFromRunes(t *testing.T) {

	tests := []struct{
		Value []rune
	}{
		{
			Value: []rune(""),
		},



		{
			Value: []rune("apple"),
		},
		{
			Value: []rune("banana"),
		},
		{
			Value: []rune("cherry"),
		},



		{
			Value: []rune("apple banana cherry"),
		},



		{
			Value: []rune("😏 💀👻👽 😊 😍 🙂🙁"),
		},
	}

	{
		const numRand = 20
		for i:=0; i<numRand; i++ {
			test := struct{
				Value []rune
			}{
				Value: []rune(fmt.Sprintf("%d", randomness.Int63())),
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		var decoder Decoder

		const targetKey = "target_key"

		decoder.data = map[string][]interface{}{}
		decoder.data[targetKey] = []interface{}{ test.Value }

		x, err := decoder.String(targetKey)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := []rune(x)

		if expected, actual := len(test.Value), len(y); expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
		for index:=0; index<len(y);index++ {
			if expected, actual := test.Value[index], y[index]; expected != actual {
				t.Errorf("For test #%d and index %d, expected %v, but actually got %v.", testNumber, index, expected, actual)
				continue
			}
		}
	}
}

func TestDecoderStringFromString(t *testing.T) {

	tests := []struct{
		Value string
	}{
		{
			Value: "",
		},



		{
			Value: "apple",
		},
		{
			Value: "banana",
		},
		{
			Value: "cherry",
		},


		{
			Value: "apple banana cherry",
		},



		{
			Value: "ONE",
		},
		{
			Value: "TWO",
		},
		{
			Value: "THREE",
		},



		{
			Value: "ONE TWO THREE",
		},



		{
			Value: "😏 💀👻👽 😊 😍 🙂🙁",
		},
	}

	{
		const numRand = 20
		for i:=0; i<numRand; i++ {
			test := struct{
				Value string
			}{
				Value: fmt.Sprintf("%d", randomness.Int63()),
			}
			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		var decoder Decoder

		const targetKey = "target_key"

		decoder.data = map[string][]interface{}{}
		decoder.data[targetKey] = []interface{}{ test.Value }

		x, err := decoder.String(targetKey)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := string(x)

		if expected, actual := test.Value, y; expected != actual {
			t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
			continue
		}
	}
}
