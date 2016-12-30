package fiatdriver

type Registrar interface {
	Register(string, Driver) error
	Obtain(string) (Driver, error)
}
