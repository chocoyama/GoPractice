package main

import (
    "time"
    "log"
)

func f(msg string) {
    log.Println(msg)
}

func main() {
    go f("hello")
    go func() {
        log.Println("gopher")
    }()
    time.Sleep(time.Second) // goroutineが他で実行されていてもmain()が終わるとプロセスが終了する
}
