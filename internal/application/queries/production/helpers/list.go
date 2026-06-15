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
		growingcycle.StatusPlanned:    "Запланирован",
		growingcycle.StatusActive:     "Активен",
		growingcycle.StatusPaused:     "Приостановлен",
		growingcycle.StatusHarvesting: "Сбор урожая",
		growingcycle.StatusCompleted:  "Завершен",
		growingcycle.StatusFailed:     "Неудача",
		growingcycle.StatusArchived:   "Архивирован",
	}
	stages := map[growingcycle.CycleStage]string{
		growingcycle.StagePlanning:    "Планирование",
		growingcycle.StageGermination: "Прорастание",
		growingcycle.StageSeedling:    "Рассада",
		growingcycle.StageVegetative:  "Вегетация",
		growingcycle.StageFlowering:   "Цветение",
		growingcycle.StageFruiting:    "Плодоношение",
		growingcycle.StageHarvesting:  "Сбор урожая",
		growingcycle.StageCompleted:   "Завершено",
	}

	return HelperResponse{
		Statuses: statuses,
		Stages:   stages,
		Methods: map[growingcycle.ProductionMethod]string{
			growingcycle.ProductionMethodSeedling:  "Рассадный",
			growingcycle.ProductionMethodDirectSow: "Прямой посев",
		},
	}, nil
}
