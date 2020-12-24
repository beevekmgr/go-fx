package main

import "go.uber.org/fx"

func main() {
	fx.New(
		bundlefx.Module,
	).Run()
}
