package automationrule

import (
	env "github.com/samurenkoroma/agro-platform/internal/domain/environment/aggregate/sensor"
)

type Operator string

const (
	Less           Operator = "LESS"
	LessOrEqual    Operator = "LESS_OR_EQUAL"
	Greater        Operator = "GREATER"
	GreaterOrEqual Operator = "GREATER_OR_EQUAL"
	Equal          Operator = "EQUAL"
)

type Trigger struct {
	SensorType env.Type
	Operator   Operator
	Value      float64
}
