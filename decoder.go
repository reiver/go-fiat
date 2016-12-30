package fiat

import (
	"github.com/reiver/go-fiat/driver"

	"fmt"
	"time"
)

// Decoder is a tool that enables you to get the inferred key-value pairs from a struct.
//
// Example:
//
//	type OrderCreated struct {
//		eventName    fiat.Const `fiat.name:"name"    fiat.value:"ORDER_CREATED"`
//		eventVersion fiat.Const `fiat.name:"version" fiat.value:"1.3.2"`
//	
//		eventData flat.Eval `fiat.name:"data" fiat.type="json" fiat.value="name,version,price"`
//	
//		Item  string `fiat:"item"`
//		Price string `fiat.name:"price"`
//	}
//	
//	var orderCreated OrderCreated
//	
//	orderCreated.Item  = "Pizza"
//	orderCreated.Price = "$2.50"
//
//	decoder := Decoder{Value: orderCreated}
//	
//	name, err := decoder.String("name")
//	if nil != err {
//		return err
//	}
//	
//	version, err := decoder.String("version")
//	if nil != err {
//		return err
//	}
//	
//	data, err := decoder.String("data")
//	if nil != err {
//		return err
//	}
type Decoder struct {
	Value interface{}

	err error
	data map[string][]interface{}
}

func (receiver *Decoder) init() {
	if nil == receiver {
		return
	}

	value := receiver.Value

	m, err := extract(value)
	if nil != err {
		receiver.err = err
		return
	}
	if nil == m {
		receiver.err = errInternalError
		return
	}

	receiver.data = m
}


func (receiver *Decoder) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	if err := receiver.err; nil != err {
		return err
	}

	data := receiver.data
	if nil == data {
		receiver.init()
		if err := receiver.err; nil != err {
			return err
		}
		data = receiver.data
		if nil == data {
			return errInternalError
		}
	}

	return nil
}

func (receiver *Decoder) Bool(name string) (bool, error) {
	if nil == receiver {
		return false, errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return false, err
	}

	switch value := valueInterface.(type) {
	case bool:
		return bool(value), nil
	default:
		return false, internalWrongTypeComplainer{expectedType:"bool", actualType:fmt.Sprintf("%T",value)}
	}
}

func (receiver *Decoder) Float32(name string) (float32, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return 0, err
	}

	switch value := valueInterface.(type) {
	case float32:
		return float32(value), nil
	case uint8:
		return float32(value), nil
	case uint16:
		return float32(value), nil
	case int8:
		return float32(value), nil
	case int16:
		return float32(value), nil
	default:
		return 0, internalWrongTypeComplainer{expectedType:"float32", actualType:fmt.Sprintf("%T",value)}
	}
}

func (receiver *Decoder) Float64(name string) (float64, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return 0, err
	}

	switch value := valueInterface.(type) {
	case float32:
		return float64(value), nil
	case float64:
		return float64(value), nil
	case uint8:
		return float64(value), nil
	case uint16:
		return float64(value), nil
	case uint32:
		return float64(value), nil
	case int8:
		return float64(value), nil
	case int16:
		return float64(value), nil
	case int32:
		return float64(value), nil
	default:
		return 0, internalWrongTypeComplainer{expectedType:"float64", actualType:fmt.Sprintf("%T",value)}
	}
}

func (receiver *Decoder) Float64x2(name string) ([2]float64, error) {
	if nil == receiver {
		return [2]float64{0.0,0.0}, errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return [2]float64{0.0,0.0}, err
	}

	switch value := valueInterface.(type) {
	case [2]float64:
		return [2]float64(value), nil
	default:
		return [2]float64{0.0,0.0}, internalWrongTypeComplainer{expectedType:"[2]float64", actualType:fmt.Sprintf("%T",value)}
	}
}

func (receiver *Decoder) Int(name string) (int, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return 0, err
	}

	switch value := valueInterface.(type) {
	case uint8:
		return int(value), nil
	case uint16:
		return int(value), nil
	case int:
		return int(value), nil
	case int8:
		return int(value), nil
	case int16:
		return int(value), nil
	case int32:
		return int(value), nil
	default:
		return 0, internalWrongTypeComplainer{expectedType:"int", actualType:fmt.Sprintf("%T",value)}
	}
}

func (receiver *Decoder) Int8(name string) (int8, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return 0, err
	}

	switch value := valueInterface.(type) {
	case int8:
		return int8(value), nil
	default:
		return 0, internalWrongTypeComplainer{expectedType:"int8", actualType:fmt.Sprintf("%T",value)}
	}
}

func (receiver *Decoder) Int16(name string) (int16, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return 0, err
	}

	switch value := valueInterface.(type) {
	case uint8:
		return int16(value), nil
	case int8:
		return int16(value), nil
	case int16:
		return int16(value), nil
	default:
		return 0, internalWrongTypeComplainer{expectedType:"int16", actualType:fmt.Sprintf("%T",value)}
	}
}

func (receiver *Decoder) Int32(name string) (int32, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return 0, err
	}

	switch value := valueInterface.(type) {
	case uint8:
		return int32(value), nil
	case uint16:
		return int32(value), nil
	case int8:
		return int32(value), nil
	case int16:
		return int32(value), nil
	case int32:
		return int32(value), nil
	default:
		return 0, internalWrongTypeComplainer{expectedType:"int32", actualType:fmt.Sprintf("%T",value)}
	}
}

func (receiver *Decoder) Int64(name string) (int64, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return 0, err
	}

	switch value := valueInterface.(type) {
	case uint8:
		return int64(value), nil
	case uint16:
		return int64(value), nil
	case uint32:
		return int64(value), nil
	case int:
		return int64(value), nil
	case int8:
		return int64(value), nil
	case int16:
		return int64(value), nil
	case int32:
		return int64(value), nil
	case int64:
		return int64(value), nil
	default:
		return 0, internalWrongTypeComplainer{expectedType:"int64", actualType:fmt.Sprintf("%T",value)}
	}
}

func (receiver *Decoder) Interface(name string) (interface{}, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	if err := receiver.err; nil != err {
		return 0, err
	}

	data := receiver.data
	if nil == data {
		receiver.init()
		if err := receiver.err; nil != err {
			return 0, err
		}
		if nil == data {
			return 0, errInternalError
		}
	}

	values, ok := data[name]
	if !ok {
		return 0, internalNotFoundComplainer{name:name}
	}
	if nil == values {
		return 0, internalNotFoundComplainer{name:name}
	}
	if 0 >= len(values) {
		return 0, internalNotFoundComplainer{name:name}
	}

	valueInterface := values[0]
	if nil == valueInterface {
		return 0, internalNotFoundComplainer{name:name}
	}

	if command, ok := valueInterface.(internalCommandWrapper); ok {

		driver, err := fiatdriver.Registry.Obtain(command.Type)
		if nil != err {
			return 0, err
		}

		value, err := driver.Eval(command.Code, data)
		if nil != err {
			return 0, err
		}

		valueInterface = value
	}

	return valueInterface, nil
}

func (receiver *Decoder) Nada(name string) (struct{}, error) {
	if nil == receiver {
		return struct{}{}, errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return struct{}{}, err
	}

	switch value := valueInterface.(type) {
	case struct{}:
		return struct{}(value), nil
	default:
		return struct{}{}, internalWrongTypeComplainer{expectedType:"struct{}", actualType:fmt.Sprintf("%T",value)}
	}
}

func (receiver *Decoder) String(name string) (string, error) {
	if nil == receiver {
		return "", errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return "", err
	}

	switch value := valueInterface.(type) {
	case []byte:
		return string(value), nil
	case rune:
		return string(value), nil
	case []rune:
		return string(value), nil
	case string:
		return string(value), nil
	default:
		return "", internalWrongTypeComplainer{expectedType:"string", actualType:fmt.Sprintf("%T",value)}
	}
}

func (receiver *Decoder) Time(name string) (time.Time, error) {
	if nil == receiver {
		return time.Time{}, errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return time.Time{}, err
	}

	switch value := valueInterface.(type) {
	case time.Time:
		return time.Time(value), nil
	default:
		return time.Time{}, internalWrongTypeComplainer{expectedType:"time.Time", actualType:fmt.Sprintf("%T",value)}
	}
}

func (receiver *Decoder) Uint8(name string) (uint8, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return 0, err
	}

	switch value := valueInterface.(type) {
	case uint8:
		return uint8(value), nil
	default:
		return 0, internalWrongTypeComplainer{expectedType:"uint8", actualType:fmt.Sprintf("%T",value)}
	}
}

func (receiver *Decoder) Uint16(name string) (uint16, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return 0, err
	}

	switch value := valueInterface.(type) {
	case uint8:
		return uint16(value), nil
	case uint16:
		return uint16(value), nil
	default:
		return 0, internalWrongTypeComplainer{expectedType:"uint16", actualType:fmt.Sprintf("%T",value)}
	}
}

func (receiver *Decoder) Uint32(name string) (uint32, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return 0, err
	}

	switch value := valueInterface.(type) {
	case uint8:
		return uint32(value), nil
	case uint16:
		return uint32(value), nil
	case uint32:
		return uint32(value), nil
	default:
		return 0, internalWrongTypeComplainer{expectedType:"uint32", actualType:fmt.Sprintf("%T",value)}
	}
}

func (receiver *Decoder) Uint64(name string) (uint64, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}


	valueInterface, err := receiver.Interface(name)
	if nil != err {
		return 0, err
	}

	switch value := valueInterface.(type) {
	case uint:
		return uint64(value), nil
	case uint8:
		return uint64(value), nil
	case uint16:
		return uint64(value), nil
	case uint32:
		return uint64(value), nil
	case uint64:
		return uint64(value), nil
	default:
		return 0, internalWrongTypeComplainer{expectedType:"uint64", actualType:fmt.Sprintf("%T",value)}
	}
}
