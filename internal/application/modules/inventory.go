package modules

import (
	itemCmd "github.com/samurenkoroma/agro-platform/internal/application/commands/inventory/item"
	warehouseCmd "github.com/samurenkoroma/agro-platform/internal/application/commands/inventory/warehouse"
	itemQuery "github.com/samurenkoroma/agro-platform/internal/application/queries/inventory/item"
	movQuery "github.com/samurenkoroma/agro-platform/internal/application/queries/inventory/movement"
	warehouseQuery "github.com/samurenkoroma/agro-platform/internal/application/queries/inventory/warehouse"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	itemProj "github.com/samurenkoroma/agro-platform/internal/infrastructure/projection/postgres/inventory/item"
	movProj "github.com/samurenkoroma/agro-platform/internal/infrastructure/projection/postgres/inventory/movement"
	warehouseProj "github.com/samurenkoroma/agro-platform/internal/infrastructure/projection/postgres/inventory/warehouse"
	"github.com/samurenkoroma/agro-platform/pkg/utils"
)

func MakeInventoryModule(uow uow.UnitOfWork, db uow.DB) Module {
	ih := itemCmd.NewItemHandler(uow)
	wh := warehouseCmd.NewWarehouseHandler(uow)

	return Module{
		Commands: []*CommandCNF{
			{
				RouteName: "inventory.create_item",
				Handler:   ih.Create,
				Decoder:   utils.DecodeJSON[itemCmd.CreateItemCommand],
			},
			{
				RouteName: "inventory.receive",
				Handler:   ih.Receive,
				Decoder:   utils.DecodeJSON[itemCmd.ReceiveCommand],
			},
			{
				RouteName: "inventory.reserve",
				Handler:   ih.Reserve,
				Decoder:   utils.DecodeJSON[itemCmd.ReserveCommand],
			},
			{
				RouteName: "inventory.consume",
				Handler:   ih.Consume,
				Decoder:   utils.DecodeJSON[itemCmd.ConsumeCommand],
			},
			{
				RouteName: "inventory.mark_lost",
				Handler:   ih.MarkLost,
				Decoder:   utils.DecodeJSON[itemCmd.MarkLostCommand],
			},
			{
				RouteName: "inventory.transfer",
				Handler:   ih.Transfer,
				Decoder:   utils.DecodeJSON[itemCmd.TransferCommand],
			},
			{
				RouteName: "inventory.create_warehouse",
				Handler:   wh.Create,
				Decoder:   utils.DecodeJSON[warehouseCmd.CreateWarehouseCommand],
			},
			{
				RouteName: "inventory.archive_warehouse",
				Handler:   wh.Archive,
				Decoder:   utils.DecodeJSON[warehouseCmd.ArchiveWarehouseCommand],
			},
		},
		Queries: []*QueryCNF{
			{
				RouteName: "inventory.get_item",
				Handler:   itemQuery.NewGetOne(itemProj.New(db)),
				Decoder:   utils.DecodeJSON[itemQuery.GetOneQuery],
			},
			{
				RouteName: "inventory.list_items",
				Handler:   itemQuery.NewList(itemProj.New(db)),
				Decoder:   utils.DecodeJSON[itemQuery.ListQuery],
			},
			{
				RouteName: "inventory.list_movements",
				Handler:   movQuery.NewList(movProj.New(db)),
				Decoder:   utils.DecodeJSON[movQuery.ListQuery],
			},
			{
				RouteName: "inventory.list_warehouses",
				Handler:   warehouseQuery.NewList(warehouseProj.New(db)),
				Decoder:   utils.DecodeJSON[warehouseQuery.ListQuery],
			},
		},
	}
}
