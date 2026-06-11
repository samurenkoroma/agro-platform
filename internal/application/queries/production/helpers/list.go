package helpers

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	growingcycle "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"
)

type helperHandler struct {
}

func New() queries.Handler {
	return &helperHandler{}
}

type HelperQuery struct {
}

type HelperResponse struct {
	Statuses map[growingcycle.CycleStatus]string      `json:"statuses"`
	Stages   map[growingcycle.CycleStage]string       `json:"stages"`
	Methods  map[growingcycle.ProductionMethod]string `json:"methods"`
}

func (h *helperHandler) Ask(ctx context.Context, payload any) (any, error) {
	statuses := map[growingcycle.CycleStatus]string{
		growingcycle.StatusPlanned:    "запланирован",
		growingcycle.StatusActive:     "активен",
		growingcycle.StatusPaused:     "приостановлен",
		growingcycle.StatusHarvesting: "сбор урожая",
		growingcycle.StatusCompleted:  "завершен",
		growingcycle.StatusFailed:     "неудача",
		growingcycle.StatusArchived:   "архивирован",
	}
	stages := map[growingcycle.CycleStage]string{
		growingcycle.StagePlanning:    "планирование",
		growingcycle.StageGermination: "прорастание",
		growingcycle.StageSeedling:    "рассада",
		growingcycle.StageVegetative:  "вегетация",
		growingcycle.StageFlowering:   "цветение",
		growingcycle.StageFruiting:    "плодоношение",
		growingcycle.StageHarvesting:  "сбор урожая",
		growingcycle.StageCompleted:   "завершено",
	}

	return HelperResponse{
		Statuses: statuses,
		Stages:   stages,
		Methods: map[growingcycle.ProductionMethod]string{
			growingcycle.ProductionMethodSeedling:  "рассадный",
			growingcycle.ProductionMethodDirectSow: "прямой посев",
		},
	}, nil
}
