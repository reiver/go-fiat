package fiat

// Eval is used to add code that is evaluated to obtain a value, to a Go struct.
//
// For example:
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
type Eval struct{}
