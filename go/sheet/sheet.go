package sheet

import (
    "Tumai/go/file_roller"
    "log"
    "fmt"
    "strconv"

    "github.com/xuri/excelize/v2"
)


var file_path string
var index int
var sheet string

type Reme struct {
    Name string             //  col 1
    Phonetic string         //  col 2
    Meaning string          //  col 3
    Note string             //  col 4
    Mastered bool           //  col 5
}

type Scoreboard struct {
    All int
    Reviewed int
    Correct int
    Wrong int
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


func Save_sheet (input []Reme) bool {
    log.Println (file_path, len(input))

    if !file_roller.File_good (file_path) {
        return false
    }

    f, err := excelize.OpenFile (file_path)
    if err != nil {
        log.Printf ("Failed to open sheet for writing: %v\n%v\n", file_path, err)
    }
    defer f.Close ()

    f.SetActiveSheet (0)

    f.SetCellValue (sheet, fmt.Sprintf ("A1"), "Name")
    f.SetCellValue (sheet, fmt.Sprintf ("B1"), "Phonetic")
    f.SetCellValue (sheet, fmt.Sprintf ("C1"), "Meaning")
    f.SetCellValue (sheet, fmt.Sprintf ("D1"), "Notes")
    f.SetCellValue (sheet, fmt.Sprintf ("E1"), "Mastered")

    f.SetColWidth (sheet, "A", "B", 20)
    f.SetColWidth (sheet, "C", "D", 50)
    f.SetColWidth (sheet, "E", "E", 10)

    for i, r := range input {
        row := i + 2
        f.SetCellValue (sheet, fmt.Sprintf ("A%d", row), r.Name)
        f.SetCellValue (sheet, fmt.Sprintf ("B%d", row), r.Phonetic)
        f.SetCellValue (sheet, fmt.Sprintf ("C%d", row), r.Meaning)
        f.SetCellValue (sheet, fmt.Sprintf ("D%d", row), r.Note)
        f.SetCellValue (sheet, fmt.Sprintf ("E%d", row), strconv.FormatBool (r.Mastered))
    }

    save_err := f.SaveAs (file_path)
    if save_err != nil {
        log.Printf ("Failed to save sheet => %v\n%v\n", file_path, save_err)
        return false
    }
    return true
}


func Open () (bool, [][]string) {
    f, err := excelize.OpenFile (file_path)
    if err != nil {
        log.Printf ("Could not open file: %s\n%s\n", file_path, err)
        return false, nil
    }
    defer f.Close ()

    for _, name := range f.GetSheetList () {
        sheet = name
        break
    }
    rows, err := f.GetRows (sheet)
    if err != nil {
        log.Printf ("Could not read file: %s\n%s\n", file_path, err)
        return false, nil
    }

    for i, val := range rows {
        if i < 1 && len (val) < 2 {
            return false, nil
        }
    }

    return true, rows[1:]
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



//  TODO => shufler