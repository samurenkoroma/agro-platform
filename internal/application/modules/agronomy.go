package modules

import (
	"github.com/samurenkoroma/agro-platform/internal/application/commands/agronomy/crop"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/agronomy/season"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/agronomy/variety"
	getcrop "github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/crop/get_crop"
	listcrops "github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/crop/list_crops"
	"github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/season/list_seasons"
	getvariety "github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/variety/get_variety"
	listvarieties "github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/variety/list_varieties"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/pkg/utils"
)

func MakeAgronomyModule(uow uow.UnitOfWork, db uow.DB) Module {
	return Module{
		Commands: []*CommandCNF{
			{
				RouteName: "agronomy.create_crop",
				Handler:   crop.NewHandler(uow).Create,
				Decoder:   utils.DecodeJSON[crop.CreateCropCommand],
			},
			{
				RouteName: "agronomy.create_variety",
				Handler:   variety.NewHandler(uow).Create,
				Decoder:   utils.DecodeJSON[variety.CreateVarietyCommand],
			},
			{
				RouteName: "agronomy.create_season",
				Handler:   season.NewHandler(uow).Create,
				Decoder:   utils.DecodeJSON[season.CreateSeasonCmd],
			},
		},
		Queries: []*QueryCNF{
			{
				RouteName: "agronomy.get_crop",
				Handler:   getcrop.New(db),
				Decoder:   utils.DecodeJSON[getcrop.Query],
			},
			{
				RouteName: "agronomy.list_crops",
				Handler:   listcrops.New(db),
				Decoder:   utils.DecodeJSON[listcrops.Query],
			},
			{
				RouteName: "agronomy.get_variety",
				Handler:   getvariety.New(db),
				Decoder:   utils.DecodeJSON[getvariety.Query],
			},
			{
				RouteName: "agronomy.list_varieties",
				Handler:   listvarieties.New(db),
				Decoder:   utils.DecodeJSON[listvarieties.Query],
			},
			{
				RouteName: "agronomy.list_seasons",
				Handler:   listseasons.New(db),
				Decoder:   utils.DecodeJSON[listseasons.Query],
			},
		},
	}
}
