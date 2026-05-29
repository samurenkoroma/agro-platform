package modules

import (
	"github.com/samurenkoroma/agro-platform/internal/application/commands"
	"github.com/samurenkoroma/agro-platform/internal/application/queries"
)

type CommandCNF struct {
	RouteName string
	Handler   commands.Handler
	Decoder   commands.DecoderFunc
}
type QueryCNF struct {
	RouteName string
	Handler   queries.Handler
	Decoder   queries.DecoderFunc
}

type Module struct {
	Commands []*CommandCNF
	Queries  []*QueryCNF
}

func (f *Module) RegisterCommands(router commands.Router) {
	for _, cmd := range f.Commands {
		router.Register(cmd.RouteName, cmd.Handler, cmd.Decoder)
	}
}

func (f *Module) RegisterQueries(router queries.Router) {
	for _, q := range f.Queries {
		router.Register(q.RouteName, q.Handler, q.Decoder)
	}
}
