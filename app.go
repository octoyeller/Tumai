package main

import (
	"context"
	"fmt"

    "github.com/wailsapp/wails/v2/pkg/runtime"

	"Tumai/go"
)

// App struct
type App struct {
	ctx context.Context

	list []file_roller.Entry
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx


    a.list = file_roller.Read_list ()

}


func (a *App) Exit_app () {
    fmt.Println ("Quit!")
    // TODO save
    runtime.Quit (a.ctx)
}

func (a *App) Get_list () (list []file_roller.Entry) {
    if len (a.list) < 1 {
        a.list = file_roller.Read_list ()
    }
    return a.list
}


func (a *App) AddList () () {}

func (a *App) RemoveList () () {}

func (a *App) SaveList () {
    if !file_roller.Save_list (a.list) {
        fmt.Println ("Could not save Tumai file at:", file_roller.List_file_path)
    }
}