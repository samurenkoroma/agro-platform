package provider

func Deps(name ProviderName, inMemory bool) ProviderDeps {
	return ProviderDeps{
		Name:     name,
		InMemory: inMemory,
	}
}

type ProviderName string

type ProviderDeps struct {
	Name     ProviderName
	InMemory bool
}
