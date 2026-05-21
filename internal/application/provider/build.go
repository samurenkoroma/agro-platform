package provider

func Deps(name ProviderName, inMemory bool) ProviderDeps {
	return ProviderDeps{
		Name:     name,
		InMemory: inMemory,
	}
}

func InMemoryDeps(name ProviderName) ProviderDeps {
	return ProviderDeps{
		Name:     name,
		InMemory: true,
	}
}

type ProviderName string

type ProviderDeps struct {
	Name     ProviderName
	InMemory bool
}
