package fiat

import (
	"github.com/reiver/go-fiat/driver"

	"github.com/reiver/go-cast"

	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

var (
	targetNameKeyShort  = nameKeyShort
	targetNameKeyNormal = nameKeyNormal
	targetNameKeyLong   = "fiat.target.name"
)

// decode loads values into the struct referred to with `v`, with the
// data in `data`.
//
// For example:
//
//	data := map[string][]interface{}{
//		"ONE":   1,
//		"TWO":   2.0,
//		"THREE": "three",
//		"FieldOne":    true,
//		"FIELD_TWO":   -7.0,
//		"field_three": 33,
//		"field FOUR":  "APPLE BANANA CHERRY",
//	}
//	
//	type TheStruct struct {
//		FieldOne   bool
//		FieldTwo   float64 `fiat:"FIELD_TWO"`
//		FieldThree int64   `fiat.name:"field_three"`
//		FieldFounr string  `fiat.target.name:"field FOUR"`
//	}
//	
//	var x TheStruct
//	
//	err := decode(data, &x)
//	
//	// Result:
//	//
//	// x.FieldOne   == true
//	// x.FieldTwo   == -7.0
//	// x.FieldThree == 33
//	// x.FieldFounr == "APPLE BANANA CHERRY"
func decode(data map[string][]interface{}, v interface{}) error {

	var reflectedStructValue reflect.Value
	var reflectedStructType  reflect.Type
	{
		// This needs to get at the struct.
		reflectedValue := reflect.ValueOf(v)
		reflectedType  := reflect.TypeOf(v)
		if nil == reflectedType {
			return errNilReflectedType
		}
		for reflect.Ptr == reflectedValue.Kind() {
			reflectedValue = reflectedValue.Elem()
			reflectedType = reflectedType.Elem()
			if nil == reflectedType {
				return errNilReflectedType
			}
		}
		if reflect.Struct != reflectedType.Kind() {
			return fmt.Errorf("Unsupported Type: %T", v)
		}

		reflectedStructValue = reflectedValue
		reflectedStructType  = reflectedType
	}

	numFields := reflectedStructType.NumField()
	for fieldNumber:=0; fieldNumber<numFields; fieldNumber++ {

		reflectedFieldValue := reflectedStructValue.Field(fieldNumber)
		if !reflectedFieldValue.CanSet() {
			continue
		}

		name := targetStructFieldSrcName(fieldNumber, reflectedStructValue)

		values, ok := data[name]
		if !ok {
			continue
		}
		if 1 > len(values) {
			continue
		}
		value := values[0]


		if command, ok := value.(internalCommandWrapper); ok {

			driver, err := fiatdriver.Registry.Obtain(command.Type)
			if nil != err {
				if _, ok := err.(fiatdriver.NotFoundComplainer); ok {
					var buffer bytes.Buffer

					fmt.Fprintf(&buffer, "Could not find driver for fiat.type=%q for field with fiat.name=%q.", command.Type, name)
					if "" == command.Type {
						fmt.Fprintf(&buffer, " Is it possible that someone forgot to add a \"fiat.type\" struct tag to the field?")
					}

					err = errors.New(buffer.String())
				}

				return err
			}

			evaledValue, err := driver.Eval(command.Code, data)
			if nil != err {
				return err
			}

			value = evaledValue
		}


		var castedValue interface{}
		{
			switch reflectedFieldValue.Kind() {
			//case reflect.Array:
			//@TODO
			case                reflect.Bool:
				casted, err := cast.Bool(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Complex64:
				casted, err := cast.Complex64(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Complex128:
				casted, err := cast.Complex128(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Float32:
				casted, err := cast.Float32(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Float64:
				casted, err := cast.Float64(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Int:
				casted, err := cast.Int(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Int8:
				casted, err := cast.Int8(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Int16:
				casted, err := cast.Int16(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Int32:
				casted, err := cast.Int32(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Int64:
				casted, err := cast.Int64(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.String:
				casted, err := cast.String(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Uint:
				casted, err := cast.Uint(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Uint8:
				casted, err := cast.Uint8(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Uint16:
				casted, err := cast.Uint16(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Uint32:
				casted, err := cast.Uint32(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			case                reflect.Uint64:
				casted, err := cast.Uint64(value)
				if nil != err {
					err = errCouldNotSet(casted, name, err)
					return err
				}
				castedValue = casted
			default:
				{
					pointerToReflectedFieldValue := reflectedFieldValue.Addr()
					switch x := pointerToReflectedFieldValue.Interface().(type) {
					case sql.Scanner:
						if err := x.Scan(value); nil != err {
							return err
						}
						continue
					default:
						err := fmt.Errorf("Cannot cast into something of type %T, for struct field name %q.", x, name)
						return err
					}
				}
			}
		}


		if err := func() (err error) {
			defer func() {
				if r := recover(); nil != r {
					err = fmt.Errorf("Could not set value ([%T] %v) for struct field named %q because: (%T) %v", castedValue, castedValue, name, r, r)
				}
			}()
			reflectedFieldValue.Set( reflect.ValueOf(castedValue) )
			return nil
		}(); nil != err {
			return err
		}
	}

	return nil
}

// targetStructFieldSrcName figures out what name should be used for a struct field, when trying
// to inject a value into a struct field.
//
// For example, if we have the struct:
//
//	type TheStruct struct {
//		FieldOne   bool
//		FieldTwo   float64 `fiat:"FIELD_TWO"`
//		FieldThree int64   `fiat.name:"field_three"`
//		FieldFounr string  `fiat.target.name:"field FOUR"`
//	}
//
// ... then the names we would get are:
//
// TheStruct.FieldOne   -> "FieldOne"
//
// TheStruct.FieldTwo   -> "FIELD_TWO"
//
// TheStruct.FieldThree -> "field_three"
//
// TheStruct.FieldFour  -> "field FOUR"
func targetStructFieldSrcName(fieldNumber int, reflectedStructValue reflect.Value) string {

	// Initialize.
	reflectedStructField := reflectedStructValue.Type().Field(fieldNumber)

	// Figure out the "name" the user wants to use.
	name, ok := reflectedStructField.Tag.Lookup(targetNameKeyLong)
	if !ok {
		name, ok = reflectedStructField.Tag.Lookup(targetNameKeyNormal)
		if !ok {
			name, ok = reflectedStructField.Tag.Lookup(targetNameKeyShort)
			if !ok {
				name = reflectedStructField.Name
			}
		}
	}

	return name
}
