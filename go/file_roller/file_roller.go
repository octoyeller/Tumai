package file_roller

import (
    "fmt"
    "os"
    "log"
    "path/filepath"
    "errors"
    "bufio"
    "strconv"
    "time"
    "strings"
    "sort"

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
        line += "0"
    }

    line = date + "\t" + e.Path

    return line
}

func Sort_list(list []Entry)  []Entry {

    sort.Slice(list[:], func(i, j int) bool {
      return list[i].Date_accessed > list[j].Date_accessed
    })

    return list
}

var List_file_path string



func Make_list () (list []Entry) {
    file_contents := fetch_file ()

    for _, line := range file_contents {
        list = append (list, make_Entry (line))
    }

    return list
}


func Read_list (l_in []Entry) (l_out []Entry) {

    if len (l_in) < 1 {
        l_in = Make_list ()
    }

    l_out = Sort_list (l_in)

    return l_out
}



func Add_list (path string, e_in []Entry) (e_out []Entry) {
    if !File_good (path) {
        fmt.Println ("Invalid file?")
        return e_in
    }
    name := filepath.Base (path)

    _, date := Get_time ()

    var e Entry
    e.Name = name
    e.Date_accessed = date
    e.Path = path

    e_out = append (e_out, e)
    for _, val := range e_in {
        e_out = append (e_out, val)
    }

    return e_out
}


func Get_time () (ts string, ti int) {

    now := time.Now()
    y := now.Year()
    m := int(now.Month())
    d := now.Day()
    h := now.Hour()
    mm := now.Minute()
    s := now.Second()


    var ys, ms, ds, hs, mms, ss string

    var err error
    ys = strconv.Itoa (y)
    ms = strconv.Itoa (m)
    if len (ms) < 2 {
        temp := ms
        ms = "0" + temp
    }
    ds = strconv.Itoa (d)
    if len (ds) < 2 {
        temp := ds
        ds = "0" + temp
    }

    hs = strconv.Itoa (h)
    if len (hs) < 2 {
        temp := hs
        hs = "0" + temp
    }

    mms = strconv.Itoa (mm)
    if len (mms) < 2 {
        temp := mms
        mms = "0" + temp
    }
    ss = strconv.Itoa (s)
    if len (ms) < 2 {
        temp := ss
        ss = "0" + temp
    }


    ts = ys + ms + ds + hs + mms + ss
    ti, err = strconv.Atoi (ts)
    if err != nil {
        ti = 10000000000000
    }

    return ts, ti
}

func Remove_list () {}




func Save_list (list []Entry) (successful bool) {
    file_contents := []string{}
    for _, obj := range list {
        file_contents = append (file_contents, obj.export())
    }

    file, err := os.Create(List_file_path)
    if err != nil {
        log.Println ("Save problem =>", err)
        return false
    }
    defer file.Close ()

    for _, val := range file_contents {
        _, err = file.WriteString (val + "\n")
        if err != nil {
            log.Println ("Save problem =>", err)
            return false
        }
    }

    log.Println ("List saved.")
    return true
}




func fetch_file () (file_contents []string) {

    home_dir, err := os.UserHomeDir ()
    if err != nil {
        log.Fatal ("Getting env ->", err)
    }

    list_path := filepath.Join (home_dir, ".Tumai.list")


    if !File_good (list_path) {
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


func File_good (path string) bool {

    file, err := os.Stat (path)

    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            log.Println ("File does not exist:", path)
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

