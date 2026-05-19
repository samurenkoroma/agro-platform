package startgrowingcycle

import (
	"context"

	gc "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"

	sharedrepo "github.com/samurenkoroma/agro-platform/internal/shared/repository"

	tx "github.com/samurenkoroma/agro-platform/internal/shared/tx"

	eb "github.com/samurenkoroma/agro-platform/internal/shared/bus/event"
)

type Handler struct {
	repositories sharedrepo.Provider
	uowManager   tx.Manager
	eventBus     eb.Bus
}

func NewHandler(repositories sharedrepo.Provider, uowManager tx.Manager, eventBus eb.Bus) *Handler {
	return &Handler{
		repositories: repositories,
		uowManager:   uowManager,
		eventBus:     eventBus,
	}
}

func (h *Handler) Handle(ctx context.Context, cmd Command) error {
	transaction, err := h.uowManager.Begin(ctx)

	if err != nil {
		return err
	}

	uow := tx.NewUnitOfWork(transaction)

	defer uow.Rollback(ctx)

	unit, err := h.repositories.
		Spatial().
		ProductionUnits().
		GetByID(cmd.ProductionUnitID)

	if err != nil {
		return err
	}

	if unit == nil {
		return ErrProductionUnitNotFound
	}

	active, err := h.repositories.
		Production().
		GrowingCycles().
		GetActiveByUnit(cmd.ProductionUnitID)

	if err != nil {
		return err
	}

	if len(active) > 0 {
		return ErrCycleAlreadyExists
	}

	cycle :=
		gc.New(
			cmd.FarmID,
			cmd.CropID,
			cmd.VarietyID,
			cmd.ProductionUnitID,
		)

	if cmd.ExpectedHarvestAt != nil {
		cycle.SetExpectedHarvest(*cmd.ExpectedHarvestAt)
	}

	err = h.repositories.
		Production().
		GrowingCycles().
		Save(cycle)

	if err != nil {
		return err
	}

	uow.AddEvents(cycle.Events()...)

	err = uow.Commit(ctx)

	if err != nil {
		return err
	}

	return h.eventBus.Publish(
		ctx,
		uow.Events()...,
	)
}
