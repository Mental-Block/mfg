package cli

import (
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/danielgtaylor/huma/v2/humacli"
)

func New(OnStart func(), onStop func()) humacli.CLI {
	return humacli.New(func(hooks humacli.Hooks, options *struct{}) {
		hooks.OnStart(OnStart)
		hooks.OnStop(onStop)
	})
}
