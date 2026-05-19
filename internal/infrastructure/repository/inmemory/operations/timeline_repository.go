package repository

import "github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"

type timeLineRepository struct {
}

func NewTimeLineRepository() repository.TimeLineRepository {
	return &timeLineRepository{}
}
