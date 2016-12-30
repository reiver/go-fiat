package fiat

import (
	"testing"
)

func TestExtractStruct(t *testing.T) {

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

		someJSON Eval `fiat.type:"json" fiat.value:"server_name,Rating"`
	}

	var foodOrder FoodOrder

	foodOrder.Type       = "cherry pie"
	foodOrder.ServerName = "Jane Doe"
	foodOrder.Rating     = 5

	data, err := extract(foodOrder)
	if nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}


	if expected, actual := 9, len(data); expected != actual {
		t.Errorf("Expected %d, but actually got %d.", expected, actual)
		return
	}


	{
		const expectedKey   = "event_name"
		const expectedValue = "FOOD_REQUESTED"

		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(string)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
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

		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(string)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
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

		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(string)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
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

		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(string)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
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

		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(string)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
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

		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(int)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
			return
		}

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		const expectedKey    = "sql"
		      expectedValue :=
`INSERT INTO food_order_events
SET name    = $1
  , version = $2
  , data    = $3
`
		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(string)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
			return
		}

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		const expectedKey    = "doc"
		      expectedValue :=
`To be, or not to be, that is the question.
The fox knows many things, but the hedgehog knows one big thing.
`
		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(string)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
			return
		}

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		const expectedKey    = "someJSON"
		      expectedValue := internalCommandWrapper{
				Type: "json",
				Code: "server_name,Rating",
			}

		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(internalCommandWrapper)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
			return
		}

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}
}

func TestExtractStructPtr(t *testing.T) {

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

		someJSON Eval `fiat.type:"json" fiat.value:"server_name,Rating"`
	}

	var foodOrder FoodOrder

	foodOrder.Type       = "cherry pie"
	foodOrder.ServerName = "Jane Doe"
	foodOrder.Rating     = 5

	data, err := extract(&foodOrder)
	if nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}


	if expected, actual := 9, len(data); expected != actual {
		t.Errorf("Expected %d, but actually got %d.", expected, actual)
		return
	}


	{
		const expectedKey   = "event_name"
		const expectedValue = "FOOD_REQUESTED"

		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(string)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
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

		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(string)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
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

		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(string)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
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

		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(string)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
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

		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(string)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
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

		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(int)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
			return
		}

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		const expectedKey    = "sql"
		      expectedValue :=
`INSERT INTO food_order_events
SET name    = $1
  , version = $2
  , data    = $3
`
		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(string)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
			return
		}

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		const expectedKey    = "doc"
		      expectedValue :=
`To be, or not to be, that is the question.
The fox knows many things, but the hedgehog knows one big thing.
`
		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(string)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
			return
		}

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}

	{
		const expectedKey    = "someJSON"
		      expectedValue := internalCommandWrapper{
				Type: "json",
				Code: "server_name,Rating",
			}

		actualValueSliceOfInterface, ok := data[expectedKey]
		if !ok {
			t.Errorf("Expected value for key %q, but actually was not there. (%t)", expectedKey, ok)
			return
		}

		if expected, actual := 1, len(actualValueSliceOfInterface); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		actualValueInterface := actualValueSliceOfInterface[0]

		actualValue, ok := actualValueInterface.(internalCommandWrapper)
		if !ok {
			t.Errorf("Expected value to have a different type, but actually was %T.", actualValueInterface)
			return
		}

		if expected, actual := expectedValue, actualValue; expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}
}
