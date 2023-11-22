package main

import (
	"astroboy/internal/dependencies"
	"astroboy/internal/router"
	"astroboy/internal/webhandler"
)

func main() {
	deps := dependencies.Init()
	wh := webhandler.NewWebHandler(deps)
	r := router.New(wh)

	r.Logger.Fatal(r.Start(":" + deps.Config.WebPort))
}
