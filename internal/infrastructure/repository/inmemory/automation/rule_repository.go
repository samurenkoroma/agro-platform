package repository

import "github.com/samurenkoroma/agro-platform/internal/domain/automation/repository"

type ruleRepository struct {
}

func NewRuleRepository() repository.RuleRepository {
	return &ruleRepository{}
}
