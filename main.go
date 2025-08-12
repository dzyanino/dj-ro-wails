package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"

	"dj-ro/internal/services"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	randomizer := services.NewRandomizer()
	dijkstra := services.NewDijkstra()
	longest_path := services.NewLongestPath()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "RO-Dijkstra",
		Width:     1280,
		Height:    720,
		MinWidth:  960,
		MinHeight: 540,

		AssetServer: &assetserver.Options{
			Assets: assets,
		},

		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 64},

		LogLevel: logger.WARNING,

		Linux: &linux.Options{
			ProgramName:         "Ro-Dijkstra",
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
			WindowIsTranslucent: true,
		},

		OnStartup: app.startup,
		Bind: []interface{}{
			app,
			randomizer,
			dijkstra,
			longest_path,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
