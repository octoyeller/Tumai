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
	irem int
	sb sheet.Scoreboard
	what bool
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

//    a.list = file_roller.Read_list (a.list)

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

func (a *App) RemoveList (i int) () {
    a.list = append (a.list[:i], a.list[i+1:]...)
}

func (a *App) SaveList () {
    if !file_roller.Save_list (a.list) {
        log.Println ("Could not save Tumai file at:", file_roller.List_file_path)
    }
}




func (a *App) Run_file (i int) {
    _, ti := file_roller.Get_time ()
    a.list[i].Date_accessed = ti
    a.reme = sheet.Run_sheet (a.list[i])
    a.irem = 0
    a.sb.All = len (a.reme)
    a.Update_score ()
    a.what = false
}


func (a *App) Get_score () (sheet.Scoreboard) {
    a.Update_score ()
    return a.sb
}

func (a *App) Update_score () {
    a.sb.Reviewed = a.irem
    a.sb.Correct = 0
    a.sb.Wrong = 0
    for _, val := range a.reme[:a.irem] {
        if val.Mastered {
            a.sb.Correct++
        }
    }
    a.sb.Wrong = a.sb.Reviewed - a.sb.Correct
}



func (a *App) Get_card () (sheet.Reme) {
    return a.reme[a.irem]
}


func (a *App) RidSheet () {
    a.reme = nil
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
    sheet.Save_sheet (a.reme)
}


func (a *App) Next_card () (sheet.Reme) {
    a.reme[a.irem].Mastered = false
    if a.irem < len (a.reme) {
        a.irem++
    }
    return a.reme[a.irem]
}

func (a *App) Learn_card () (sheet.Reme) {
    a.reme[a.irem].Mastered = true
    if a.irem < len (a.reme) {
        a.irem++
    }
    return a.reme[a.irem]
}

func (a *App) Previous_card () (sheet.Reme) {
    if a.irem > 0 {
        a.irem--
    }
    return a.reme[a.irem]

}


func (a *App) SwitchWhat () bool {
    a.what = !a.what
    return a.what
}

func (a *App) GetWhat () bool {
    return a.what
}


func (a *App) SwitchSelect () {
    fmt.Println ("Select")
}


func (a *App) SwitchOrder () {
    fmt.Println ("Oreder")
}