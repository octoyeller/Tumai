package file_roller

import (
    "fmt"
    "os"
    "log"
    "path/filepath"
    "errors"
    "bufio"
    "strconv"
    "strings"

    "github.com/wneessen/go-fileperm"
)


type Entry struct {

    Path string
    Date_accessed int
    Name string

}

func (e *Entry) export () (line string) {
    date := strconv.Itoa (e.Date_accessed)
    zeros := len (date) - 14
    for i := 0; i < zeros; i++ {
        date += "0"
    }

    line = date + "\t" + e.Path

    return line
}

var List_file_path string



func Read_list () (list []Entry) {
    file_contents := fetch_file ()

    for _, line := range file_contents {
        list = append (list, make_Entry (line))
    }


    return list
}



func Add_list () {}

func Remove_list () {}




func Save_list (list []Entry) (successful bool) {
    file_contents := []string{}
    for _, obj := range list {
        file_contents = append (file_contents, obj.export())
    }

    file, err := os.Create(List_file_path)
    if err != nil {
        fmt.Println ("Save problem =>", err)
        return false
    }
    defer file.Close ()

    for _, val := range file_contents {
        _, err = file.WriteString (val + "\n")
        if err != nil {
            fmt.Println ("Save problem =>", err)
            return false
        }
    }

    fmt.Println ("List saved.")
    return true
}




func fetch_file () (file_contents []string) {

    home_dir, err := os.UserHomeDir ()
    if err != nil {
        log.Fatal ("Getting env ->", err)
    }

    list_path := filepath.Join (home_dir, ".Tumai.list")


    if !file_good (list_path) {
        log.Fatal ("Needed file is unavailable:", list_path)
    }




    file, err := os.Open (list_path)
    if err != nil && !errors.Is (err, os.ErrNotExist) {
        log.Fatal ("Failed to open file: ", list_path)
    }
    defer file.Close ()


    scanner := bufio.NewScanner (file)
    for scanner.Scan () {
        file_contents = append (file_contents, scanner.Text ())
    }

    List_file_path = list_path
    return file_contents

}


func file_good (path string) bool {

    file, err := os.Stat (path)

    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            fmt.Println ("User file does not exist, will create it:", path)
            return true
        }
        return false
    }

    if file.IsDir () {
        return false
    }


    perm, err := fileperm.New (path)

    if !perm.UserReadable () {
        return false
    }

    if !perm.UserWritable () {
        return false
    }

    return true
}



func make_Entry (input string) (e Entry) {

    offset := strings.Index (input, "\t")
    if offset < 0 {
        e.Name = "Nuh-uh"
        e.Path = ""
        e.Date_accessed = 0
        return e
    }

    e.Path = input[offset+1:]
    e.Name = filepath.Base (e.Path)

    date := input[:offset]
    date_o, err := strconv.Atoi (date)
    e.Date_accessed = date_o
    if err != nil {
        e.Date_accessed = 0
    }

    return e
}

