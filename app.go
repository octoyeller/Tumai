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


    a.list = file_roller.Read_list (a.list)

}


func (a *App) Exit_app () {
    fmt.Println ("Quit!")
    // TODO save
    runtime.Quit (a.ctx)
}

func (a *App) Get_list () (list []file_roller.Entry) {

    a.list = file_roller.Read_list (a.list)

    return a.list
}


func (a *App) AddList () () {
    path, err := a.Pick_file ()
    if err != nil {
        fmt.Println (err)
    }
    if path == "" {
        fmt.Println ("No file chosen.")
        return
    }

    fmt.Println ("Picking file:", path)
    a.list = file_roller.Add_list (path, a.list)
}

func (a *App) RemoveList () () {}

func (a *App) SaveList () {
    if !file_roller.Save_list (a.list) {
        fmt.Println ("Could not save Tumai file at:", file_roller.List_file_path)
    }
}




func (a *App) Run_file (i int) {
    fmt.Println (i)
    _, ti := file_roller.Get_time ()
    a.list[i].Date_accessed = ti
}







func (a *App) Pick_file() (string, error) {
    path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
        Title: "Select a file",
        // optional filters:
        Filters: []runtime.FileFilter{
            {DisplayName: "All files", Pattern: "*"},
            {DisplayName: "ODF Spreadsheet (*.ods)", Pattern: "*.ods"},
            {DisplayName: "Excel (*.xlsx)", Pattern: "*.xlsx"},
            {DisplayName: "Excel '97 (*.xls)", Pattern: "*.xls"},
        },
    })
    if err != nil {
        return "", err
    }
    return path, nil // empty string if user cancelled
}