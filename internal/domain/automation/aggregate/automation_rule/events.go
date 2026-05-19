package automationrule

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventRuleCreated  = "automation.rule.created"
	EventRuleEnabled  = "automation.rule.enabled"
	EventRuleDisabled = "automation.rule.disabled"
	EventRuleExecuted = "automation.rule.executed"
	EventRuleArchived = "automation.rule.archived"
)

type RuleCreated struct {
	ev.BaseEvent
}

func NewRuleCreated(id vo.ID) RuleCreated {
	return RuleCreated{ev.NewBaseEvent(id, EventRuleCreated)}
}

type RuleEnabled struct {
	ev.BaseEvent
}

func NewRuleEnabled(id vo.ID) RuleEnabled {
	return RuleEnabled{ev.NewBaseEvent(id, EventRuleEnabled)}
}

type RuleDisabled struct {
	ev.BaseEvent
}

func NewRuleDisabled(id vo.ID) RuleDisabled {
	return RuleDisabled{ev.NewBaseEvent(id, EventRuleDisabled)}
}

type RuleExecuted struct {
	ev.BaseEvent
}

func NewRuleExecuted(id vo.ID) RuleExecuted {
	return RuleExecuted{ev.NewBaseEvent(id, EventRuleExecuted)}
}

type RuleArchived struct {
	ev.BaseEvent
}

func NewRuleArchived(id vo.ID) RuleArchived {
	return RuleArchived{ev.NewBaseEvent(id, EventRuleArchived)}
}
