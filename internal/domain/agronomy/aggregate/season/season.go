package season

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type SeasonStatus string

const (
	StatusPlanning  SeasonStatus = "planning"
	StatusCurrent   SeasonStatus = "current"
	StatusCompleted SeasonStatus = "completed"
)

// Season - агрономический сезон
type Season struct {
	ev.BaseAggregate
	ID        vo.ID
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Status    SeasonStatus

	CreatedBy vo.ID
	OwnerID   vo.ID

	CreatedAt  time.Time
	UpdatedAt  time.Time
	ArchivedAt *time.Time
}

func New(
	name string,
	startDate, endDate time.Time,
	status SeasonStatus,
	createdBy vo.ID,
	ownerId vo.ID,
) (*Season, error) {
	if name == "" {
		return nil, ErrInvalidName
	}
	if status == StatusPlanning && endDate.Before(time.Now()) {
		return nil, ErrInvalidPlanningInPast
	}
	if startDate.After(endDate) {
		return nil, ErrInvalidPeriod
	}
	if createdBy == "" {
		return nil, ErrInvalidCreatedBy
	}
	now := time.Now()

	s := &Season{
		ID:        vo.NewID(),
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    status,
		CreatedBy: createdBy,
		OwnerID:   ownerId,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return s, nil
}

// Activate активирует сезон
func (s *Season) Activate() error {
	if s.Status != StatusPlanning {
		return ErrInvalidStatusTransition
	}

	now := time.Now()
	if now.Before(s.StartDate) {
		return ErrSeasonNotStarted
	}

	s.Status = StatusCurrent

	s.AddEvent(SeasonActivated{
		SeasonID: s.ID.String(),
		Name:     s.Name,
	})

	return nil
}

// Complete завершает сезон
func (s *Season) Complete() error {
	if s.Status != StatusCurrent {
		return ErrInvalidStatusTransition
	}

	s.Status = StatusCompleted

	s.AddEvent(SeasonCompleted{
		SeasonID: s.ID.String(),
		Name:     s.Name,
	})

	return nil
}

// IsActive проверяет, активен ли сезон в указанную дату
func (s *Season) IsActiveAt(date time.Time) bool {
	return !date.Before(s.StartDate) && !date.After(s.EndDate)
}

func (s *Season) Archived() {
	now := time.Now()
	s.ArchivedAt = &now

	s.AddEvent(SeasonArchived{
		SeasonID: s.ID.String(),
	})
}
