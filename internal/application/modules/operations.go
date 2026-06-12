package modules

import (
	opCmd "github.com/samurenkoroma/agro-platform/internal/application/commands/operations/operation_event"
	taskCmd "github.com/samurenkoroma/agro-platform/internal/application/commands/operations/task"
	opQuery "github.com/samurenkoroma/agro-platform/internal/application/queries/operations/operation_event"
	taskQuery "github.com/samurenkoroma/agro-platform/internal/application/queries/operations/task"
	tlQuery "github.com/samurenkoroma/agro-platform/internal/application/queries/operations/timeline"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	opProjection "github.com/samurenkoroma/agro-platform/internal/infrastructure/projection/postgres/operations/operation_event"
	taskProjection "github.com/samurenkoroma/agro-platform/internal/infrastructure/projection/postgres/operations/task"
	tlProjection "github.com/samurenkoroma/agro-platform/internal/infrastructure/projection/postgres/operations/timeline"
	"github.com/samurenkoroma/agro-platform/pkg/utils"
)

func MakeOperationsModule(uow uow.UnitOfWork, db uow.DB) Module {
	return Module{
		Commands: []*CommandCNF{
			{
				RouteName: "operations.create_task",
				Handler:   taskCmd.NewTaskHandler(uow).Create,
				Decoder:   utils.DecodeJSON[taskCmd.CreateTaskCommand],
			},
			{
				RouteName: "operations.assign_task",
				Handler:   taskCmd.NewTaskHandler(uow).Assign,
				Decoder:   utils.DecodeJSON[taskCmd.AssignTaskCommand],
			},
			{
				RouteName: "operations.start_task",
				Handler:   taskCmd.NewTaskHandler(uow).Start,
				Decoder:   utils.DecodeJSON[taskCmd.TaskIDCommand],
			},
			{
				RouteName: "operations.complete_task",
				Handler:   taskCmd.NewTaskHandler(uow).Complete,
				Decoder:   utils.DecodeJSON[taskCmd.TaskIDCommand],
			},
			{
				RouteName: "operations.cancel_task",
				Handler:   taskCmd.NewTaskHandler(uow).Cancel,
				Decoder:   utils.DecodeJSON[taskCmd.TaskIDCommand],
			},
			{
				RouteName: "operations.record_operation",
				Handler:   opCmd.NewOperationHandler(uow).Record,
				Decoder:   utils.DecodeJSON[opCmd.RecordOperationCommand],
			},
		},
		Queries: []*QueryCNF{
			{
				RouteName: "operations.get_task",
				Handler:   taskQuery.NewGetOne(taskProjection.New(db)),
				Decoder:   utils.DecodeJSON[taskQuery.GetOneQuery],
			},
			{
				RouteName: "operations.list_tasks",
				Handler:   taskQuery.NewList(taskProjection.New(db)),
				Decoder:   utils.DecodeJSON[taskQuery.ListQuery],
			},
			{
				RouteName: "operations.get_timeline",
				Handler:   tlQuery.NewGet(tlProjection.New(db)),
				Decoder:   utils.DecodeJSON[tlQuery.GetQuery],
			},
			{
				RouteName: "operations.list_operations",
				Handler:   opQuery.NewList(opProjection.New(db)),
				Decoder:   utils.DecodeJSON[opQuery.ListQuery],
			},
		},
	}
}
