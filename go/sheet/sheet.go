package sheet

import (
    "Tumai/go/file_roller"
    "log"
    "fmt"
    "strconv"
    "math/rand"
    "strings"
    "path/filepath"
    "regexp"
    "os"

    "github.com/xuri/excelize/v2"
    odso "github.com/fwilhe2/rechenbrett"    // write
    odsi "github.com/knieriem/odf/ods"       // read
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
        log.Printf ("Failed to read cards from: %v\n", file_path)
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

    ext := strings.ToLower (filepath.Ext (file_path))
    if ext ==".ods" {
        return saveODS (input)
    }
    return saveXLSX (input)
}


func saveXLSX (input []Reme) bool {
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


func saveODS (input []Reme) bool {
    h := make([]odso.Cell, 5)
    h[0] = odso.MakeCell ("Name", "string")
    h[1] = odso.MakeCell ("Phonetic", "string")
    h[2] = odso.MakeCell ("Meaning", "string")
    h[3] = odso.MakeCell ("Notes", "string")
    h[4] = odso.MakeCell ("Mastered", "string")

    var tbl [][]odso.Cell
    tbl = append (tbl, h)

    for _, val := range input {
        row := make ([]odso.Cell, 5)
        row[0] = odso.MakeCell (val.Name, "string")
        row[1] = odso.MakeCell (val.Phonetic, "string")
        row[2] = odso.MakeCell (val.Meaning, "string")
        row[3] = odso.MakeCell (val.Note, "string")
        row[4] = odso.MakeCell (strconv.FormatBool(val.Mastered), "string")
        tbl = append (tbl, row)
    }

    doc, err := odso.MakeSpreadsheetWithName ("Sheet1", tbl)
    if err != nil {
        log.Printf ("Could not create a spreadsheet: %v\n", err)
        return false
    }

    s, save_err := odso.MakeOds (doc)
    if save_err != nil {
        log.Printf ("Failed to make ODS: %v\n", save_err)
        return false
    }

    wrt_err := os.WriteFile(file_path, s.Bytes (), 0644)
    if wrt_err != nil {
        log.Printf("Failed to save ODS file: %v\n", wrt_err)
        return false
    }

    return true
}


func Open () (bool, [][]string) {
    ext := strings.ToLower (filepath.Ext (file_path) )
    if ext == ".ods" {
        return openODS ()
    }
    return openXLSX ()
}




func openXLSX () (bool, [][]string) {
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



func openODS () (bool, [][]string) {
    f, err := odsi.Open (file_path)
    if err != nil {
        log.Printf ("Could not open file: %s\n%s\n", file_path, err)
        return false, nil
    }
    defer f.Close ()

    var doc odsi.Doc
    if err := f.ParseContent (&doc); err != nil {
        log.Printf ("Error reading file: %s\n%s\n", file_path, err)
        return false, nil
    }

    if len (doc.Table) == 0 {
        log.Printf ("No tables in the file: %s\n", file_path)
        return false, nil
    }


    reee := regexp.MustCompile (`<[^>]*>`)

    var rows [][]string
    for _, row := range doc.Table[0].Row {
        var row_content []string

        for _, cell := range row.Cell {
            var cell_text []string
            for _, p := range cell.P {
                text := reee.ReplaceAllString (p.XML, "")
                if text != "" {
                    cell_text = append (cell_text, text)
                }
            }

            rape_text := strings.Join (cell_text, "\n")

            repeat := cell.RepeatedCols

            if repeat < 1 {
                repeat = 1
            }

            if rape_text == "" && repeat > 50 {
                repeat = 1
            }

            for i := 0; i < repeat; i++ {
                row_content = append (row_content, rape_text)
            }

        }

        if len (row_content[0]) < 1 {
            continue
        }
        rows = append (rows, row_content)
    }

    if len (rows) < 2 {
        return false, nil
    }

    return true, rows[1:]
}



func Get_reme (rows [][]string) (ra []Reme) {

    r := Reme{}
    var i int

    for _, val := range rows {
        i = len (val)
        if i < 3 {
            continue
        }

        r.Name = val[0]
        r.Phonetic = val[1]
        r.Meaning = val[2]

        if i > 4 {
            if (val[4] != "true") {
                r.Mastered = false
            } else {
                r.Mastered = true
            }
            r.Note = val[3]

        } else if i > 3 {
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



func Shuffler (in []Reme) (out []int) {

    ii := len (in)
    var ran int

    for i := 0; i < ii; i++ {
        ran = rand.Int () % (ii - i)
        out = append (out, ran)
        in = append (in[:ran], in[ran:]...)
    }

    return out
}


//  TODO    images?