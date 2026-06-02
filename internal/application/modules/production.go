package modules

import (
	allocationCmd "github.com/samurenkoroma/agro-platform/internal/application/commands/production/allocation"
	growingcycleCmd "github.com/samurenkoroma/agro-platform/internal/application/commands/production/growing_cycle"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/production/harvest"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/production/planting"
	allocationQuery "github.com/samurenkoroma/agro-platform/internal/application/queries/production/allocation"
	growingcycleQuery "github.com/samurenkoroma/agro-platform/internal/application/queries/production/growing_cycle"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/projection/postgres/production/allocation"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/projection/postgres/production/growing_cycle"
	"github.com/samurenkoroma/agro-platform/pkg/utils"
)

func MakeProductionModule(uow uow.UnitOfWork, db uow.DB) Module {
	return Module{
		Commands: []*CommandCNF{
			{
				RouteName: "production.create_cycle",
				Handler:   growingcycleCmd.NewGrowingCycleHandler(uow).Create,
				Decoder:   utils.DecodeJSON[growingcycleCmd.CreateCommand],
			},
			{
				RouteName: "production.allocate_production_unit",
				Handler:   allocationCmd.NewAllocationHandler(uow).AllocateProductionUnit,
				Decoder:   utils.DecodeJSON[allocationCmd.AllocateProductionUnitCommand],
			},
			{
				RouteName: "production.change_allocation",
				Handler:   allocationCmd.NewAllocationHandler(uow).Change,
				Decoder:   utils.DecodeJSON[allocationCmd.ChangeAllocationCommand],
			},
			{
				RouteName: "production.release_allocation",
				Handler:   allocationCmd.NewAllocationHandler(uow).Release,
				Decoder:   utils.DecodeJSON[allocationCmd.ReleaseAllocationCommand],
			},
			{
				RouteName: "production.planting_register",
				Handler:   planting.NewPlantingHandler(uow).Register,
				Decoder:   utils.DecodeJSON[planting.RegisterPlantingCommand],
			},
			{
				RouteName: "production.planting_change",
				Handler:   planting.NewPlantingHandler(uow).Change,
				Decoder:   utils.DecodeJSON[planting.ChangePlantingCommand],
			},
			{
				RouteName: "production.harvest_register",
				Handler:   harvest.NewHarvestHandler(uow).Register,
				Decoder:   utils.DecodeJSON[harvest.RegisterHarvestCommand],
			},
			{
				RouteName: "production.harvest_change",
				Handler:   harvest.NewHarvestHandler(uow).Change,
				Decoder:   utils.DecodeJSON[harvest.ChangeHarvestCommand],
			},
		},
		Queries: []*QueryCNF{
			{
				RouteName: "production.get_growing_cycle",
				Handler:   growingcycleQuery.NewGetOne(growingcycle.New(db)),
				Decoder:   utils.DecodeJSON[growingcycleQuery.GetOneQuery],
			},
			{
				RouteName: "production.list_growing_cycles",
				Handler:   growingcycleQuery.NewList(growingcycle.New(db)),
				Decoder:   utils.DecodeJSON[growingcycleQuery.ListQuery],
			},
			{
				RouteName: "production.summary_growing_cycles",
				Handler:   growingcycleQuery.NewSummaryHandler(growingcycle.New(db)),
				Decoder:   utils.DecodeJSON[growingcycleQuery.SummaryQuery],
			},
			{
				RouteName: "production.get_allocation",
				Handler:   allocationQuery.NewGetOne(allocation.New(db)),
				Decoder:   utils.DecodeJSON[allocationQuery.GetOneQuery],
			},
			{
				RouteName: "production.list_allocations",
				Handler:   allocationQuery.NewList(allocation.New(db)),
				Decoder:   utils.DecodeJSON[allocationQuery.ListQuery],
			},
		},
	}
}
