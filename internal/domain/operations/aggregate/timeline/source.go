package timeline

type Source string

const (
	TaskSource       Source = "TASK"
	OperationSource  Source = "OPERATION"
	HarvestSource    Source = "HARVEST"
	YieldSource      Source = "YIELD"
	TelemetrySource  Source = "TELEMETRY"
	AlertSource      Source = "ALERT"
	AutomationSource Source = "AUTOMATION"
)
