package fiat

import (
	"testing"
)

func TestDecoderDecode(t *testing.T) {

	type sourceStruct struct {
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

		ONE_a bool
		ONE_b int64
		ONE_c string

		TWO_a bool
		TWO_b int64
		TWO_c string

		THREE_a bool
		THREE_b int64
		THREE_c string

		ShortM1  string `fiat:"m1_short"`
		LongM1   string `fiat:"m1_long"`
		NormalM1 string `fiat:"m1_normal"`

		ShortM2  string `fiat.name:"m2_short"`
		LongM2   string `fiat.name:"m2_long"`
		NormalM2 string `fiat.name:"m2_normal"`

		ShortM3  string `fiat:"m3_short"`
		LongM3   string `fiat:"m3_long"`
		NormalM3 string `fiat:"m3_normal"`
	}

	var src sourceStruct

	src.ShouldBeFalse = false
	src.ShouldBeTrue  = true

	src.ShouldBe0 = int64(0)
	src.ShouldBe1 = int64(1)
	src.ShouldBe2 = int64(2)
	src.ShouldBe3 = int64(3)
	src.ShouldBe4 = int64(4)
	src.ShouldBe5 = int64(5)

	src.ShouldBeApple  = "Apple"
	src.ShouldBeBanana = "Banana"
	src.ShouldBeCherry = "Cherry"

	src.ONE_a = true
	src.ONE_b = int64(5)
	src.ONE_c = "something"

	src.TWO_a = true
	src.TWO_b = int64(10)
	src.TWO_c = "somewhere"

	src.THREE_a = true
	src.THREE_b = int64(15)
	src.THREE_c = "somehow"

	src.ShortM1  = "not this m1"
	src.NormalM1 = "not this either m1"
	src.LongM1   = "m1 THIS!"

	src.ShortM2  = "not this m2"
	src.NormalM2 = "not this either m2"
	src.LongM2   = "m2 THIS!"

	src.ShortM3  = "not this m3"
	src.NormalM3 = "not this either m3"
	src.LongM3   = "m3 THIS!"

	type destinationStruct struct {
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

		M1 string `fiat:"m1_short" fiat.name:"m1_normal" fiat.target.name:"m1_long"`
		M2 string `fiat:"m2_short" fiat.name:"m2_normal" fiat.target.name:"m2_long"`
		M3 string `fiat:"m3_short" fiat.name:"m3_normal" fiat.target.name:"m3_long"`
	}

	var dest destinationStruct

	decoder := Decoder{
		Value: &src,
	}

	if err := decoder.Decode(&dest); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	if expected, actual := false, dest.ShouldBeFalse; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := true, dest.ShouldBeTrue; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}

	if expected, actual := int64(0), dest.ShouldBe0; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := int64(1), dest.ShouldBe1; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := int64(2), dest.ShouldBe2; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := int64(3), dest.ShouldBe3; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := int64(4), dest.ShouldBe4; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := int64(5), dest.ShouldBe5; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}

	if expected, actual := "Apple", dest.ShouldBeApple; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "Banana", dest.ShouldBeBanana; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "Cherry", dest.ShouldBeCherry; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}

	if expected, actual := true, dest.F1a; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := int64(5), dest.F1b; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "something", dest.F1c; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}

	if expected, actual := true, dest.F2a; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := int64(10), dest.F2b; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "somewhere", dest.F2c; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}

	if expected, actual := true, dest.F3a; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := int64(15), dest.F3b; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "somehow", dest.F3c; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}

	if expected, actual := "m1 THIS!", dest.M1; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "m2 THIS!", dest.M2; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
	if expected, actual := "m3 THIS!", dest.M3; expected != actual {
		t.Errorf("Expected (%T) %#v, but actually got (%T) %#v.", expected, expected, actual, actual)
		return
	}
}
