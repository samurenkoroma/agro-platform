package bootstrap

import (
	"context"
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/samurenkoroma/agro-platform/internal/application/commands"
	createproductionunit "github.com/samurenkoroma/agro-platform/internal/application/commands/spatial/create_production_unit"
	"github.com/samurenkoroma/agro-platform/internal/application/queries"
	getproductionunit "github.com/samurenkoroma/agro-platform/internal/application/queries/spatial/get_production_unit"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/messaging/inmemory"
	http2 "github.com/samurenkoroma/agro-platform/internal/interfaces/http"
	configs "github.com/samurenkoroma/agro-platform/internal/shared/config"
	"github.com/samurenkoroma/agro-platform/internal/shared/tx"
	"github.com/samurenkoroma/agro-platform/pkg/utils"
)

type App struct {
	DB            *sql.DB
	CommandRouter commands.Router
	QueryRouter   queries.Router
	HTTPHandler   http.Handler
}

func Build(ctx context.Context, db *sql.DB, conf *configs.Config) (*App, error) {
	bus := inmemory.NewInMemoryEventBus()
	uowFactory := tx.NewUnitOfWorkFactory(db, bus)

	commandRouter := commands.NewRouter()
	queryRouter := queries.NewRouter()
	commandRouter.Register("spatial.create_production_unit", createproductionunit.NewProductionUnitHandler(uowFactory), utils.DecodeJSON[createproductionunit.Command])
	queryRouter.Register("GetCurrentFarm", getproductionunit.NewProductionUnitHandler(), utils.DecodeJSON[getproductionunit.GetCurrentFarmQuery])
	//bus.Register("farm.field.created", growingEventHandlers.OnFarmObjectCreated)
	//bus.Register(physicalobject.FarmObjectSchemaUpdatedEvent, growingEventHandlers.OnFarmObjectSchemaUpdated)
	//bus.Register("crop.plan.published", growingEventHandlers.OnCropPlanPublished)

	httpHandler := http2.NewRouter(http2.RouterConfig{
		CommandRouter: commandRouter,
		QueryRouter:   queryRouter,
		UowFactory:    uowFactory,
		//JWTService:    jwtService,
	})
	return &App{
		DB:            db,
		CommandRouter: commandRouter,
		QueryRouter:   queryRouter,
		HTTPHandler:   httpHandler,
	}, nil
}
