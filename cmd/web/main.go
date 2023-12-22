package main

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/router"
	"astroboy/internal/webhandlers"
)

func main() {
	deps := dependencies.Init()
	wh := webhandlers.NewWebHandler(deps)
	r := router.New(wh, deps.Config)

	r.Logger.Fatal(r.Start(":" + deps.Config.WebPort))
}
