package fiat

import (
	"math"

	"testing"
)

func TestDecoderFloat64x2FromFloat64x2(t *testing.T) {

	tests := []struct{
		Value [2]float64
	}{}
	{
		f64s := []float64{
			math.Inf(-1),
			-math.MaxFloat64,
			-math.Pi,
			-math.E,
			-math.Sqrt2,
			-1.0,
			-math.Ln2,
			-math.SmallestNonzeroFloat64,
			0.0,
			math.SmallestNonzeroFloat64,
			math.Ln2,
			1.0,
			math.Sqrt2,
			math.E,
			math.Pi,
			math.MaxFloat64,
			math.Inf(+1),

			math.NaN(),
		}

		const numRand = 20
		for i:=0; i<numRand; i++ {
			f64 := randomness.Float64()
			f64s = append(f64s, f64)

			f64  = -randomness.Float64()
			f64s = append(f64s, f64)



			f64  = randomness.Float64() * math.MaxFloat64
			f64s = append(f64s, f64)

			f64  = -randomness.Float64() * math.MaxFloat64
			f64s = append(f64s, f64)



			f64  = randomness.Float64() * 999999999
			f64s = append(f64s, f64)

			f64  = -randomness.Float64() * 999999999
			f64s = append(f64s, f64)
		}

		for _, x := range f64s {
			for _, y := range f64s {

				test := struct{
					Value [2]float64
				}{
					Value: [2]float64{x, y},
				}

				tests = append(tests, test)
			}
		}
	}


	for testNumber, test := range tests {

		var decoder Decoder

		const targetKey = "target_key"

		decoder.data = map[string][]interface{}{}
		decoder.data[targetKey] = []interface{}{ test.Value }

		x, err := decoder.Float64x2(targetKey)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		y := [2]float64(x)

		if expected, actual := test.Value[0], y[0]; expected != actual {
			if !(math.IsNaN(expected) && math.IsNaN(actual)) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
		if expected, actual := test.Value[1], y[1]; expected != actual {
			if !(math.IsNaN(expected) && math.IsNaN(actual)) {
				t.Errorf("For test #%d, expected %v, but actually got %v.", testNumber, expected, actual)
				continue
			}
		}
	}
}
