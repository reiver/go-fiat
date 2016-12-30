package fiatdriver

type Driver interface {
	Eval(code string, ctx map[string][]interface{}) (interface{}, error)
}
