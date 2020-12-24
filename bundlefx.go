package bundlefx

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

var Module = fx.Options(
	Handler.Module,
	message.Module,
	fx.Invoke(registerHooks),
)

func registerHooks(lifecycle fx.Lifecycle, h *handler.Handler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Println("Starting application in :5000")
				go h.Gin.Run(":5000")
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
