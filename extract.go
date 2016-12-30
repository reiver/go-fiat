package fiat

import (
	"fmt"
	"reflect"
)

const (
	nameKeyShort  = "fiat"
	nameKeyNormal = "fiat.name"
)

const (
	nameTypeNormal = "fiat.type"
)

const (
	valueKeyNormal    = "fiat.value"
	valueKeyMultiLine = "fiat.value.multi-line"
)

// extract infers key-value paire from a struct.
//
//
// By default, it will use the struct's field name for the key (in the inferred key-value pair).
//
// This case be overridden by using a "fiat" or "fiat.name" tag.
//
// You might want to do this, for example, if you want the name used to be snake_case.
//
// For example:
//
//	type FoodOrder struct {
//		DiaryOK bool `fiat.name:"diary_ok"`
//		MeatOK  bool `fiat:"meat_ok"`
//	}
//
// And if the values for an instance of this struct was as follows:
//
//	var foodOrder FoodOrder
//	
//	foodOrder.DiaryOK = true
//	DiaryOK.MeatOK    = false
//
// Then the key-value pairs inferred from this would be:
//
// 	"diary_ok" -> true
//	"meat_ok"  -> false
//
//
// For "normal" fields, the value (in the inferred key-value pair) will the struct's field value.
//
// However, if a struct's field is of type fiat.Const, then the "fiat.value", "fiat.value.multi-line",
// and "fiat.value.json" struct tags must be used to give it a value.
//
// Using a field of type fiat.Const enables you to, for example, have "constant" field values.
//
// NOTE: A fiat.Const field can be unexported, and still work!
func extract(v interface{}) (map[string][]interface{}, error) {

	var reflectedStructValue reflect.Value
	var reflectedStructType  reflect.Type
	{
		// This needs to get at the struct.
		reflectedValue := reflect.ValueOf(v)
		reflectedType  := reflect.TypeOf(v)
		if nil == reflectedType {
			return nil, errNilReflectedType
		}
		for reflect.Ptr == reflectedValue.Kind() {
			reflectedValue = reflectedValue.Elem()
			reflectedType = reflectedType.Elem()
			if nil == reflectedType {
				return nil, errNilReflectedType
			}
		}
		if reflect.Struct != reflectedType.Kind() {
			return nil, fmt.Errorf("Unsupported Type: %T", v)
		}

		reflectedStructValue = reflectedValue
		reflectedStructType  = reflectedType
	}

	// We use this to collect the inferred key-value pairs that we extract.
	data := map[string][]interface{}{}

	numFields := reflectedStructType.NumField()
	Loop: for fieldNumber:=0; fieldNumber<numFields; fieldNumber++ {

		reflectedStructField := reflectedStructType.Field(fieldNumber)

		var name  string
		var value interface{}
		var err   error

		switch reflect.Zero(reflectedStructField.Type).Interface().(type)  {
		case Const:
			name, value, err = extractConst(v, fieldNumber, reflectedStructValue)
			if nil != err {
				return nil, err
			}
		case Eval:
			name, value, err = extractEval(v, fieldNumber, reflectedStructValue)
			if nil != err {
				return nil, err
			}
		default:
			reflectedFieldValue := reflectedStructValue.Field(fieldNumber)
			if !reflectedFieldValue.CanInterface() {
				continue Loop
			}

			name, value, err = extractDefault(v, fieldNumber, reflectedStructValue)
			if nil != err {
				return nil, err
			}
		}

		if _, ok := data[name]; !ok {
			data[name] = []interface{}{}
		}
		data[name] = append(data[name], value)
	}

	return data, nil
}

// extractConst is used when a struct field is of the special type fiat.Const.
func extractConst(v interface{}, fieldNumber int, reflectedStructValue reflect.Value) (string, interface{}, error) {

	// Initialize.
	reflectedStructField := reflectedStructValue.Type().Field(fieldNumber)

	// Figure out the "name" the user wants to use.
	name, ok := reflectedStructField.Tag.Lookup(nameKeyShort)
	if !ok {
		name, ok = reflectedStructField.Tag.Lookup(nameKeyNormal)
		if !ok {
			name = reflectedStructField.Name
		}
	}

	// Figure out the value the user wants to use.
	if value, ok := reflectedStructField.Tag.Lookup(valueKeyNormal); ok {
		return name, value, nil
	}
	if endNeedle, ok := reflectedStructField.Tag.Lookup(valueKeyMultiLine); ok {

		value, err := extractFollowingLines(string(reflectedStructField.Tag), endNeedle)
		if nil != err {
			return "", nil, err
		}

		return name, value, nil
	}

	return "", nil, fmt.Errorf("For type %T and field %q, missing %q or %q tag on struct field of type fiat.Const. Tag: %q", v, reflectedStructField.Name, valueKeyNormal, valueKeyMultiLine, string(reflectedStructField.Tag))
}

// extractEval is used when a struct field is of the special type fiat.Eval.
func extractEval(v interface{}, fieldNumber int, reflectedStructValue reflect.Value) (string, interface{}, error) {

	// Initialize.
	reflectedStructField := reflectedStructValue.Type().Field(fieldNumber)

	// Figure out the "name" the user wants to use.
	name, ok := reflectedStructField.Tag.Lookup(nameKeyShort)
	if !ok {
		name, ok = reflectedStructField.Tag.Lookup(nameKeyNormal)
		if !ok {
			name = reflectedStructField.Name
		}
	}

	// Figure out the "tyoe" the user wants to use.
	typ, ok := reflectedStructField.Tag.Lookup(nameTypeNormal)
	if !ok {
		typ = ""
	}

	// Figure out the value the user wants to use.
	if code, ok := reflectedStructField.Tag.Lookup(valueKeyNormal); ok {
		value := internalCommandWrapper{
			Code: code,
			Type: typ,
		}

		return name, value, nil
	}
	if endNeedle, ok := reflectedStructField.Tag.Lookup(valueKeyMultiLine); ok {

		code, err := extractFollowingLines(string(reflectedStructField.Tag), endNeedle)
		if nil != err {
			return "", nil, err
		}

		value := internalCommandWrapper{
			Code: code,
			Type: typ,
		}

		return name, value, nil
	}

	return "", nil, fmt.Errorf("For type %T and field %q, missing %q or %q tag on struct field of type fiat.Eval. Tag: %q", v, reflectedStructField.Name, valueKeyNormal, valueKeyMultiLine, string(reflectedStructField.Tag))
}

// extractDefault is used when a struct field is NOT of the special type fiat.Const.
func extractDefault(v interface{}, fieldNumber int, reflectedStructValue reflect.Value) (string, interface{}, error) {

	// Initialize.
	reflectedStructField := reflectedStructValue.Type().Field(fieldNumber)

	// Figure out the "name" the user wants to use.
	name, ok := reflectedStructField.Tag.Lookup(nameKeyShort)
	if !ok {
		name, ok = reflectedStructField.Tag.Lookup(nameKeyNormal)
		if !ok {
			name = reflectedStructField.Name
		}
	}

	// Figure out the value the user wants to use.
	reflectedFieldValue := reflectedStructValue.Field(fieldNumber)
	if !reflectedFieldValue.CanInterface() {
		return "", nil, fmt.Errorf("For type %T and field %q, could not extract value.", v, reflectedStructField.Name)
	}
	value := reflectedFieldValue.Interface()

	return name, value, nil
}
