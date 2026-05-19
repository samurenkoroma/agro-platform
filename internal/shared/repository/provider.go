package repository

type Provider interface {
	Spatial() SpatialProvider

	Production() ProductionProvider

	Agronomy() AgronomyProvider

	Operations() OperationsProvider

	Inventory() InventoryProvider

	Environment() EnvironmentProvider

	Automation() AutomationProvider

	Analytics() AnalyticsProvider
}
