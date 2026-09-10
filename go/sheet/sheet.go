package sheet

import (
    "Tumai/go/file_roller"
    "log"
)


var file_path string


func Run_sheet (e file_roller.Entry) {
    log.Println (e.Name)
    file_path = e.Path
}


func Save_sheet () {
    log.Println (file_path)
}

//  TODO => get file
//  TODO => shufler
//  TODO => save
//  TODO => mark learned
//  TODO => next
//  TODO => previous