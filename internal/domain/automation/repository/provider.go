package repository

type AutomationProvider interface {
	Actuators() ActuatorRepository
	Rules() RuleRepository
}
