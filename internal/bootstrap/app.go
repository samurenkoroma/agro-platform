package bootstrap

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/samurenkoroma/agro-platform/internal/application/commands"
	createcrop "github.com/samurenkoroma/agro-platform/internal/application/commands/agronomy/create_crop"
	createvariety "github.com/samurenkoroma/agro-platform/internal/application/commands/agronomy/create_variety"
	createproductionunit "github.com/samurenkoroma/agro-platform/internal/application/commands/spatial/create_production_unit"
	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	"github.com/samurenkoroma/agro-platform/internal/application/queries/account"
	"github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/crop"
	unitOfWork "github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/jwt"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/messaging/inmemory"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/postgres"
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

	commandRouter.Register("spatial.create_production_unit", createproductionunit.NewCreateProductionUnitHandler(uow), utils.DecodeJSON[createproductionunit.Command])
	commandRouter.Register("agronomy.create_crop", createcrop.NewCreateCropHandler(uow), utils.DecodeJSON[createcrop.Command])
	commandRouter.Register("agronomy.create_variety", createvariety.NewCreateVarietyHandler(uow), utils.DecodeJSON[createvariety.Command])
	queryRouter.Register("Me", account.NewUserHandler(uow, jwtService), utils.DecodeJSON[account.MeQuery])

	queryRouter.Register("agronomy.get_crops", crop.NewCropHandler(uow), utils.DecodeJSON[crop.CropsQuery])

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
