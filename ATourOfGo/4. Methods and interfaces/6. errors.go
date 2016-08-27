package main

import (
    "fmt"
    "time"
)

/*
Goではエラーの状態をerror変数で表現する

error型はfmt.Stringerに似た組み込みのインターフェース
type error interface {
    Error() string
}

ex)
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer: ", i)
*/

type MyError struct {
    When time.Time
    What string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
    return &MyError{
        time.Now(),
        "it didn't work",
    }
}

func main() {
    if err := run(); err != nil {
        fmt.Println(err)
    }
}
