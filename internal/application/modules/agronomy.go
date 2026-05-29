package modules

import (
	createvariety "github.com/samurenkoroma/agro-platform/internal/application/commands/agronomy/create_variety"
	createcrop "github.com/samurenkoroma/agro-platform/internal/application/commands/agronomy/crop"
	getcrop "github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/crop/get_crop"
	listcrops "github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/crop/list_crops"
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
				Handler:   createcrop.NewCropHandler(uow).Create,
				Decoder:   utils.DecodeJSON[createcrop.CreateCropCommand],
			},
			{
				RouteName: "agronomy.create_variety",
				Handler:   createvariety.NewCreateVarietyHandler(uow).Create,
				Decoder:   utils.DecodeJSON[createvariety.CreateVarietyCommand],
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
		},
	}
}
