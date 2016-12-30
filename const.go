package fiat

// Const is used a constant value to a Go struct.
//
// For example:
//
//	type OrderCreated struct {
//		eventName    fiat.Const `fiat.name:"name"    fiat.value:"ORDER_CREATED"`
//		eventVersion fiat.Const `fiat.name:"version" fiat.value:"1.3.2"`
//	
//		Item  string `fiat:"item"`
//		Price string `fiat.name:"price"`
//	}
type Const struct{}
