package fiat

// Meta is used to add meta-data to a Go struct.
//
// For example:
//
//	type OrderCreated struct {
//		eventName    fiat.Meta `fiat.name:"name"    fiat.value:"ORDER_CREATED"`
//		eventVersion fiat.Meta `fiat.name:"version" fiat.value:"1.3.2"`
//	
//		Item  string `fiat:"item"`
//		Price string `fiat.name:"price"`
//	}
type Meta struct{}
