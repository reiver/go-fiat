package fiat

import (
	"fmt"

	"testing"
)

func TestDecoder(t *testing.T) {

	type FoodOrder struct {
		Name       Const  `fiat.name:"event_name"    fiat.value:"FOOD_REQUESTED"`
		Version    Const  `fiat:"event_version"      fiat.value:"1.0.0"`
		Something  Const  `                          fiat.value:"apple banana cherry"`
		Type       string `fiat:"type"`
		ServerName string `fiat.name:"server_name"`
		Rating     int

		sql        Const  `fiat.value.multi-line:""
INSERT INTO food_order_events
SET name    = $1
  , version = $2
  , data    = $3
`

		doc        Const  `fiat.value.multi-line:"{{{END}}}"
To be, or not to be, that is the question.
The fox knows many things, but the hedgehog knows one big thing.
{{{END}}}
You won't see this in the final docs.
:-)
`
		someData   Eval   `fiat.type:"json" fiat.name:"some_data" fiat.value:"server_name,type"`
	}

	var foodOrder FoodOrder

	foodOrder.Type       = "cherry pie"
	foodOrder.ServerName = "Jane Doe"
	foodOrder.Rating     = 5

	decoder:= &Decoder{Value:foodOrder}
	if err := decoder.Err(); nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}


	{
		const expectedKey   = "event_name"
		const expectedValue = "FOOD_REQUESTED"

		actualValue, err := decoder.String(expectedKey)
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		const expectedKey   = "event_version"
		const expectedValue = "1.0.0"

		actualValue, err := decoder.String(expectedKey)
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		const expectedKey   = "Something"
		const expectedValue = "apple banana cherry"

		actualValue, err := decoder.String(expectedKey)
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		const expectedKey    = "type"
		      expectedValue := foodOrder.Type

		actualValue, err := decoder.String(expectedKey)
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		const expectedKey    = "server_name"
		      expectedValue := foodOrder.ServerName

		actualValue, err := decoder.String(expectedKey)
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		const expectedKey    = "Rating"
		      expectedValue := foodOrder.Rating

		actualValue, err := decoder.Int(expectedKey)
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		const expectedKey    = "some_data"
		      expectedValue := fmt.Sprintf(`{"server_name":%q,"type":%q}`+"\n", foodOrder.ServerName, foodOrder.Type)

		actualValue, err := decoder.String(expectedKey)
		if nil != err {
			t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
			return
		}

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}
}
