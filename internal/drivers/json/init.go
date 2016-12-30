package verboten

import (
	"github.com/reiver/go-fiat/driver"
)

func init() {

	const name = "json"

	driver := internalDriver{}

	if err := fiatdriver.Registry.Register(name, driver); nil != err {
		panic(err)
	}
}
