package automationrule

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Action struct {
	ActuatorID vo.ID
	Command    string
	Payload    map[string]any
}
