package repository

type AnalyticsProvider interface {
	Metrics() MetricsRepository
	YieldRecords() YieldRecordRepository
}
