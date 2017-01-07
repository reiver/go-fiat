package fiat

// Decode loads values into the struct referred to with `v`,
// with data from fiat.Decoder.Value.
//
// For example:
//
//	decoder := fiat.Decoder{
//		Value: x,
//	}
//	
//	// ...
//	
//	type MyStruct struct {
//		Apple  bool    `fiat:"something"`
//		Banana float64 `fiat.name:"else"`
//		Cherry int64   `fiat.target.name:"here"`
//		Grape  string
//	}
//	
//	// ...
//	
//	var target MyStruct
//	
//	// ...
//	
//	err := decoder.Decode(&target)
func (receiver *Decoder) Decode(v interface{}) error {
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

	return decode(data, v)
}
