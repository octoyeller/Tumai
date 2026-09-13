package sheet

import (
    "Tumai/go/file_roller"
    "log"

    "github.com/xuri/excelize/v2"
)


var file_path string
var index int

type Reme struct {
    Name string             //  col 1
    Phonetic string         //  col 2
    Meaning string          //  col 3
    Note string             //  col 4
    Mastered bool           //  col 5
}


func Run_sheet (e file_roller.Entry) (ra []Reme) {
    index = 0
    log.Println (e.Name)
    file_path = e.Path

    success, rows := Open ()
    if !success {
        return ra
    }

    ra = Get_reme (rows)

    return ra
}


func Save_sheet () {
    log.Println (file_path)
}


func Open () (bool, [][]string) {
    f, err := excelize.OpenFile (file_path)
    if err != nil {
        log.Printf ("Could not open file: %s\n%s\n", file_path, err)
        return false, nil
    }
    defer f.Close ()

    rows, err := f.GetRows ("Sheet1")
    if err != nil {
        log.Printf ("Could not read file: %s\n%s\n", file_path, err)
        return false, nil
    }

    for _, val := range rows {
        if len (val) < 2 {
            return false, nil
        }
    }

    return true, rows
}




func Get_reme (rows [][]string) (ra []Reme) {

    r := Reme{}
    var i int

    for _, val := range rows {
        i = len (val)

        r.Name = val[0]
        r.Phonetic = val[1]
        r.Meaning = val[2]

        if i > 5 {
            if (val[4] != "true") {
                r.Mastered = true
            } else {
                r.Mastered = false
            }
            r.Note = val[3]

        } else if i > 4 {
            r.Mastered = false
            r.Note = val[3]

        } else {
            r.Mastered = false
            r.Note = ""
        }

        ra = append (ra, r)
    }

    return ra
}



//  TODO => get file
//  TODO => shufler
//  TODO => save
//  TODO => mark learned
//  TODO => next
//  TODO => previous