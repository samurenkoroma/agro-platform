package automationrule

import "errors"

var (
	ErrArchivedRule = errors.New("rule archived")
	ErrDisabledRule = errors.New("rule disabled")
)
