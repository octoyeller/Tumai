package main

import (
    // window
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	// backend
	"fmt"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

    p := fmt.Println






	// Create an instance of the app structure
	p ("Initializing window")
	app := NewApp()

	// Create application with options
	window_err := wails.Run(&options.App{
		Title:  "Tumai",   //  the window title
		Width:  750,
		Height: 550,
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

//	p ("Exiting")
}
