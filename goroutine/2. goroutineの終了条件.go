package main

import (
    "log"
    "runtime"
    "time"
)

func main() {
    go func() {
        log.Println("end") // 関数が終わる
    }()
    go func() {
        log.Println("return") // returnで抜ける
        return
    }()
    go func() {
        log.Println("exit")
        runtime.Goexit() // runtime.Goexit()を実行する
    }()
    time.Sleep(time.Second)
}
