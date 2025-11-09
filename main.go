package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	app, err := NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize app: %s", err)
	}

	err = wails.Run(&options.App{
		Title:  "krate",
		Width:  1280,
		Height: 720,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
			app.CollectionHandler, // Bind handler directly to expose all its methods
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
