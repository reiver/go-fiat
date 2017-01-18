package fiat

import (
	"testing"
)

func TestDecode(t *testing.T) {

	data := map[string][]interface{}{
		"ShouldBeFalse": []interface{}{false},
		"ShouldBeTrue":  []interface{}{true},

		"ShouldBe0": []interface{}{int64(0)},
		"ShouldBe1": []interface{}{int64(1)},
		"ShouldBe2": []interface{}{int64(2)},
		"ShouldBe3": []interface{}{int64(3)},
		"ShouldBe4": []interface{}{int64(4)},
		"ShouldBe5": []interface{}{int64(5)},

		"ShouldBeApple":  []interface{}{"Apple"},
		"ShouldBeBanana": []interface{}{"Banana"},
		"ShouldBeCherry": []interface{}{"Cherry"},

		"ONE_a": []interface{}{true},
		"ONE_b": []interface{}{int64(5)},
		"ONE_c": []interface{}{"something"},

		"TWO_a": []interface{}{true},
		"TWO_b": []interface{}{int64(10)},
		"TWO_c": []interface{}{"somewhere"},

		"THREE_a": []interface{}{true},
		"THREE_b": []interface{}{int64(15)},
		"THREE_c": []interface{}{"somehow"},

		"m1_short":  []interface{}{false},
		"m1_normal": []interface{}{false},
		"m1_long":   []interface{}{true},

		"m2_short":  []interface{}{int64(23)},
		"m2_normal": []interface{}{int64(24)},
		"m2_long":   []interface{}{int64(25)},

		"m3_short":  []interface{}{"not here"},
		"m3_normal": []interface{}{"not here either"},
		"m3_long":   []interface{}{"HERE!"},

		"sstr1": []interface{}{"Str ONE"},
		"sstr2": []interface{}{"Str TWO"},
		"sstr3": []interface{}{"Str THREE"},
	}

	type myStruct struct {
		ShouldBeFalse bool
		ShouldBeTrue  bool
		ShouldBe0 int64
		ShouldBe1 int64
		ShouldBe2 int64
		ShouldBe3 int64
		ShouldBe4 int64
		ShouldBe5 int64
		ShouldBeApple  string
		ShouldBeBanana string
		ShouldBeCherry string

		F1a bool   `fiat:"ONE_a"`
		F1b int64  `fiat:"ONE_b"`
		F1c string `fiat:"ONE_c"`

		F2a bool   `fiat.name:"TWO_a"`
		F2b int64  `fiat.name:"TWO_b"`
		F2c string `fiat.name:"TWO_c"`

		F3a bool   `fiat.target.name:"THREE_a"`
		F3b int64  `fiat.target.name:"THREE_b"`
		F3c string `fiat.target.name:"THREE_c"`

		M1 bool   `fiat:"m1_short" fiat.name:"m1_normal" fiat.target.name:"m1_long"`
		M2 int64  `fiat:"m2_short" fiat.name:"m2_normal" fiat.target.name:"m2_long"`
		M3 string `fiat:"m3_short" fiat.name:"m3_normal" fiat.target.name:"m3_long"`

		SStr1 scannableString `fiat:"sstr1"`
		SStr2 scannableString `fiat.name:"sstr2"`
		SStr3 scannableString `fiat.target.name:"sstr3"`
	}

	var x myStruct

	if err := decode(data, &x); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	if expected, actual := false, x.ShouldBeFalse; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := true, x.ShouldBeTrue; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}

	if expected, actual := int64(0), x.ShouldBe0; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := int64(1), x.ShouldBe1; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := int64(2), x.ShouldBe2; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := int64(3), x.ShouldBe3; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := int64(4), x.ShouldBe4; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := int64(5), x.ShouldBe5; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}

	if expected, actual := "Apple", x.ShouldBeApple; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := "Banana", x.ShouldBeBanana; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := "Cherry", x.ShouldBeCherry; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}

	if expected, actual := true, x.F1a; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := int64(5), x.F1b; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := "something", x.F1c; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}

	if expected, actual := true, x.F2a; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := int64(10), x.F2b; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := "somewhere", x.F2c; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}

	if expected, actual := true, x.F3a; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := int64(15), x.F3b; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := "somehow", x.F3c; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}

	if expected, actual := true, x.M1; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := int64(25), x.M2; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}
	if expected, actual := "HERE!", x.M3; expected != actual {
		t.Errorf("Expected %v, but actually got %v.", expected, actual)
		return
	}

	{
		actual, err := x.SStr1.String()
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		if expected := "Str ONE"; expected != actual {
			t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
			return
		}
	}
	{
		actual, err := x.SStr2.String()
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		if expected := "Str TWO"; expected != actual {
			t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
			return
		}
	}
	{
		actual, err := x.SStr3.String()
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		if expected := "Str THREE"; expected != actual {
			t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
			return
		}
	}
}

func TestDecodeMultipleTargetsFromSingleSource(t *testing.T) {

	const expectedValue = "SUCCESS!"

	data := map[string][]interface{}{
		"FIELD_1": []interface{}{ expectedValue },
	}

	type myStructTestDecodeMultipleTargetsFromSingleSource struct {
		FIELD_1         string
		SomeField       string `fiat:"FIELD_1"`
		AnotherField    string `fiat.name:"FIELD_1"`
		YetAnotherField string `fiat.target.name:"FIELD_1"`
		ShouldNotGetIt  string
	}

	var x myStructTestDecodeMultipleTargetsFromSingleSource

	if err := decode(data, &x); nil != err {
		t.Errorf("Expected an error, but did not actually got one: (%T) %v", err, err)
		return
	}

	if expected, actual := expectedValue, x.FIELD_1; expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}
	if expected, actual := expectedValue, x.SomeField; expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}
	if expected, actual := expectedValue, x.AnotherField; expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}
	if expected, actual := expectedValue, x.YetAnotherField; expected != actual {
		t.Errorf("Expected %q, but actually got %q.", expected, actual)
		return
	}


	if notExpected, actual := expectedValue, x.ShouldNotGetIt; notExpected == actual {
		t.Errorf("Did not Expect %q, but actually got %q.", notExpected, actual)
		return
	}
}

func TestDecodeFailCannotCastInto1(t *testing.T) {

	data := map[string][]interface{}{
		"Field1": []interface{}{ "(1,2,3,4)" },
	}

	type subStruct struct {
		x float64
		y float64
		z float64
		w float64
	}

	type myStructTestDecodeFail struct {
		Field1 subStruct
	}

	var x myStructTestDecodeFail

	if err := decode(data, &x); nil == err {
		t.Errorf("Expected an error, but did not actually got one: %v", err)
		return
	}
}

func TestDecodeFailCannotCastInto2(t *testing.T) {

	data := map[string][]interface{}{
		"field_1": []interface{}{ "(1,2,3,4)" },
	}

	type subStruct struct {
		x float64
		y float64
		z float64
		w float64
	}

	type myStructTestDecodeFail struct {
		Field1 subStruct `fiat:"field_1"`
	}

	var x myStructTestDecodeFail

	if err := decode(data, &x); nil == err {
		t.Errorf("Expected an error, but did not actually got one: %v", err)
		return
	}
}

func TestDecodeFailCannotCastInto3(t *testing.T) {

	data := map[string][]interface{}{
		"Field_1": []interface{}{ "(1,2,3,4)" },
	}

	type subStruct struct {
		x float64
		y float64
		z float64
		w float64
	}

	type myStructTestDecodeFail struct {
		Field1 subStruct `fiat.name:"Field_1"`
	}

	var x myStructTestDecodeFail

	if err := decode(data, &x); nil == err {
		t.Errorf("Expected an error, but did not actually got one: %v", err)
		return
	}
}

func TestDecodeFailCannotCastInto4(t *testing.T) {

	data := map[string][]interface{}{
		"FIELD_1": []interface{}{ "(1,2,3,4)" },
	}

	type subStruct struct {
		x float64
		y float64
		z float64
		w float64
	}

	type myStructTestDecodeFail struct {
		Field1 subStruct `fiat.target.name:"FIELD_1"`
	}

	var x myStructTestDecodeFail

	if err := decode(data, &x); nil == err {
		t.Errorf("Expected an error, but did not actually got one: %v", err)
		return
	}
}
