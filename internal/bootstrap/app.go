package bootstrap

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/samurenkoroma/agro-platform/internal/application/commands"
	"github.com/samurenkoroma/agro-platform/internal/application/commands/account/organization"
	createvariety "github.com/samurenkoroma/agro-platform/internal/application/commands/agronomy/create_variety"
	createcrop "github.com/samurenkoroma/agro-platform/internal/application/commands/agronomy/crop"
	createproductionunit "github.com/samurenkoroma/agro-platform/internal/application/commands/spatial"
	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	"github.com/samurenkoroma/agro-platform/internal/application/queries/account"
	getcrop "github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/crop/get_crop"
	listcrops "github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/crop/list_crops"
	catalog "github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/variety"
	unitOfWork "github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/jwt"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/messaging/inmemory"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/postgres"
	crop2 "github.com/samurenkoroma/agro-platform/internal/infrastructure/projection/postgres/agronomy/crop"
	http2 "github.com/samurenkoroma/agro-platform/internal/interfaces/http"
	configs "github.com/samurenkoroma/agro-platform/internal/shared/config"
	"github.com/samurenkoroma/agro-platform/pkg/utils"
)

type App struct {
	DB            unitOfWork.DB
	CommandRouter commands.Router
	QueryRouter   queries.Router
	HTTPHandler   http.Handler
}

func Build(ctx context.Context, pool *pgxpool.Pool, conf *configs.Config) (*App, error) {
	bus := inmemory.NewInMemoryEventBus()
	uow := postgres.NewUnitOfWork(ctx, pool, bus)
	jwtService := jwt.NewService(conf.Auth)
	commandRouter := commands.NewRouter()
	queryRouter := queries.NewRouter()

	commandRouter.Register("spatial.create_production_unit", createproductionunit.NewCreateProductionUnitHandler(uow).Create, utils.DecodeJSON[createproductionunit.CreateCommand])

	commandRouter.Register("account.create_organization", organization.NewOrganizationHandler(uow, jwtService).Create, utils.DecodeJSON[organization.CreateOrganizationCmd])
	commandRouter.Register("account.switch_organization", organization.NewOrganizationHandler(uow, jwtService).Switch, utils.DecodeJSON[organization.SwitchOrganizationCmd])
	queryRouter.Register("Me", account.NewUserHandler(uow, jwtService), utils.DecodeJSON[account.MeQuery])

	commandRouter.Register("agronomy.create_crop", createcrop.NewCropHandler(uow).Create, utils.DecodeJSON[createcrop.CreateCropCommand])
	commandRouter.Register("agronomy.create_variety", createvariety.NewCreateVarietyHandler(uow).Handle, utils.DecodeJSON[createvariety.CreateVarietyCommand])

	cropProjection := crop2.New(pool)
	queryRouter.Register("agronomy.get_crop", getcrop.New(cropProjection), utils.DecodeJSON[getcrop.Query])
	queryRouter.Register("agronomy.list_crops", listcrops.New(cropProjection), utils.DecodeJSON[listcrops.Query])
	queryRouter.Register("agronomy.list_varieties", catalog.NewVarietyHandler(uow), utils.DecodeJSON[catalog.VarietiesQuery])

	//bus.Register("farm.field.created", growingEventHandlers.OnFarmObjectCreated)
	//bus.Register(physicalobject.FarmObjectSchemaUpdatedEvent, growingEventHandlers.OnFarmObjectSchemaUpdated)
	//bus.Register("crop.plan.published", growingEventHandlers.OnCropPlanPublished)

	httpHandler := http2.NewRouter(http2.RouterConfig{
		CommandRouter: commandRouter,
		QueryRouter:   queryRouter,
		Uow:           uow,
		JWTService:    jwtService,
	})
	return &App{
		DB:            pool,
		CommandRouter: commandRouter,
		QueryRouter:   queryRouter,
		HTTPHandler:   httpHandler,
	}, nil
}
