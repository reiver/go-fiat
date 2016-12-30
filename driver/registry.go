package fiatdriver

var (
	Registry Registrar
)

func init() {
	Registry = new(internalRegistrar)
}
