package commands

type CommandCNF struct {
	RouteName string
	Handler   Handler
	Decoder   DecoderFunc
}

type CommandModule struct {
	Routes []*CommandCNF
}
