package main

import (
	"context"
	"fmt"
	"log"

    "github.com/wailsapp/wails/v2/pkg/runtime"

	"Tumai/go/sheet"
	"Tumai/go/file_roller"
)

// App struct
type App struct {
	ctx context.Context

	list []file_roller.Entry
	reme []sheet.Reme
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
        log.Println (err)
    }
    if path == "" {
        log.Println ("No file chosen.")
        return
    }

    log.Println ("Picking file:", path)
    a.list = file_roller.Add_list (path, a.list)
}

func (a *App) RemoveList () () {}

func (a *App) SaveList () {
    if !file_roller.Save_list (a.list) {
        log.Println ("Could not save Tumai file at:", file_roller.List_file_path)
    }
}




func (a *App) Run_file (i int) {
    _, ti := file_roller.Get_time ()
    a.list[i].Date_accessed = ti
    a.reme = sheet.Run_sheet (a.list[i])
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





func (a *App) SaveSheet () {
    sheet.Save_sheet ()
}