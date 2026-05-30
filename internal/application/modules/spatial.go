package modules

import (
	createproductionunit "github.com/samurenkoroma/agro-platform/internal/application/commands/spatial/create_production_unit"
	"github.com/samurenkoroma/agro-platform/internal/application/queries/spatial/production_unit"
	"github.com/samurenkoroma/agro-platform/internal/application/queries/spatial/production_unit/get_production_unit"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	spatial "github.com/samurenkoroma/agro-platform/internal/infrastructure/projection/postgres/spatial/production_unit"
	"github.com/samurenkoroma/agro-platform/pkg/utils"
)

func MakeSpatialModule(uow uow.UnitOfWork, db uow.DB) Module {
	return Module{
		Commands: []*CommandCNF{
			{
				RouteName: "spatial.create_production_unit",
				Handler:   createproductionunit.New(uow).Create,
				Decoder:   utils.DecodeJSON[createproductionunit.CreateCommand],
			},
			{
				RouteName: "spatial.update_production_unit",
				//Handler:   createproductionunit.New(uow).Create,
				//Decoder:   utils.DecodeJSON[createproductionunit.CreateCommand],
			},
			{
				RouteName: "spatial.archive_production_unit",
				//Handler:   createproductionunit.New(uow).Create,
				//Decoder:   utils.DecodeJSON[createproductionunit.CreateCommand],
			},
			{
				RouteName: "spatial.move_production_unit",
				//Handler:   createproductionunit.New(uow).Create,
				//Decoder:   utils.DecodeJSON[createproductionunit.CreateCommand],
			},
			{
				RouteName: "spatial.clone_production_unit",
				//Handler:   createproductionunit.New(uow).Create,
				//Decoder:   utils.DecodeJSON[createproductionunit.CreateCommand],
			},
		},
		Queries: []*QueryCNF{
			{
				RouteName: "spatial.get_production_unit",
				Handler:   getproductionunit.New(db),
				Decoder:   utils.DecodeJSON[getproductionunit.Query],
			},
			{
				RouteName: "spatial.list_production_units",
				Handler:   productionunit.NewListRoots(spatial.New(db)),
				Decoder:   utils.DecodeJSON[productionunit.Query],
			},
			{
				RouteName: "spatial.get_production_unit_tree",
			},
			{
				RouteName: "spatial.get_production_unit_children",
			},
		},
	}
}
