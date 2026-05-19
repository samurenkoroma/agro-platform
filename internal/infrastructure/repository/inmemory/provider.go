package inmemory

type Provider struct {
	spatial     *SpatialProvider
	production  *ProductionProvider
	agronomy    *AgronomyProvider
	operations  *OperationsProvider
	inventory   *InventoryProvider
	environment *EnvironmentProvider
	automation  *AutomationProvider
	analytics   *AnalyticsProvider
}

func NewProvider() *Provider {
	return &Provider{
		spatial:     NewSpatialProvider(),
		production:  NewProductionProvider(),
		agronomy:    NewAgronomyProvider(),
		operations:  NewOperationsProvider(),
		inventory:   NewInventoryProvider(),
		environment: NewEnvironmentProvider(),
		automation:  NewAutomationProvider(),
		analytics:   NewAnalyticsProvider(),
	}
}
