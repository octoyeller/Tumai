package main

import (
    // window
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {


	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	window_err := wails.Run(&options.App{
		Title:  "Tumai",   //  the window title
		Linux: &linux.Options{
		    ProgramName: "Tumai",
		},
		Width:  650,
		Height: 450,
		MinWidth: 650,
		MinHeight: 450,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{},                                  //  do not touch opacity in style.css
//		BackgroundColour: &options.RGBA{R: 30, G: 30, B: 30, A: 255},       //  mandatory for proper rendering and background :c
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if window_err != nil {
		println("Error:", window_err.Error())
	}

}
