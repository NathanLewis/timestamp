package main

import (
    "os"
    "fmt"
    "time"
    "strings"
)

func main() {
    //t, err := time.Parse(layout, "2016-02-28T21:03:46")
    msg := "Usage: timestamp <time in YYYY-MM-DDTHH:MM:SS format>|now"
    if len(os.Args) < 2 {
        fmt.Println(msg)
        os.Exit(0)
    }
    arg := os.Args[1]
    input := strings.TrimRight(arg, "\n")
    if strings.Compare(input, "-h") == 0 {
        fmt.Println(msg)
        os.Exit(0)
    }
    if strings.Compare(input, "now") == 0 {
        tm := time.Now()
        fmt.Println(tm.Unix())
        os.Exit(0)
    }
    layout := "2006-01-02T15:04:05"  // The reference time
    tm, err := time.Parse(layout, input )
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(tm.Unix())
}
