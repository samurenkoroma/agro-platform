package command

type Registry interface {
	Register(name string, handler any)

	Get(name string) (any, bool)
}
