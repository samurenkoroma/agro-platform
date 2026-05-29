package modules

import (
	"github.com/samurenkoroma/agro-platform/internal/application/commands/account/organization"
	"github.com/samurenkoroma/agro-platform/internal/application/queries/account"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/jwt"
	"github.com/samurenkoroma/agro-platform/pkg/utils"
)

func MakeAccountModule(uow uow.UnitOfWork, jwt *jwt.Service) Module {
	return Module{
		Commands: []*CommandCNF{
			{
				RouteName: "account.create_organization",
				Handler:   organization.NewOrganizationHandler(uow, jwt).Create,
				Decoder:   utils.DecodeJSON[organization.CreateOrganizationCmd],
			},
			{
				RouteName: "account.switch_organization",
				Handler:   organization.NewOrganizationHandler(uow, jwt).Switch,
				Decoder:   utils.DecodeJSON[organization.SwitchOrganizationCmd],
			},
		},
		Queries: []*QueryCNF{
			{
				RouteName: "account.me",
				Handler:   account.NewUserHandler(uow, jwt),
				Decoder:   utils.DecodeJSON[account.MeQuery],
			},
		},
	}
}
