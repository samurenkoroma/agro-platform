package account

import (
	"github.com/samurenkoroma/agro-platform/internal/application/commands"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/account/organization"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/jwt"
	"github.com/samurenkoroma/agro-platform/pkg/utils"
)

func Make(uow uow.UnitOfWork, jwt *jwt.Service) commands.CommandModule {
	return commands.CommandModule{
		Routes: []*commands.CommandCNF{
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
	}
}
