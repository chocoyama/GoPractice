package main

import (
    "time"
    "fmt"
    "log"
    "math/rand"
)

func randomTime() time.Duration {
    return time.Duration(rand.Intn(1e3)) * time.Millisecond
}

func worker(msg string) <-chan string {
    receiver := make(chan string)
    for i:=0; i<300; i++ {
        go func(i int) {
            time.Sleep(randomTime())
            msg := fmt.Sprintf("%d %s done", i, msg)
            receiver <- msg
        }(i)
    }
    return receiver
}

func main() {
    receiver := worker("job")
    for {
        select {
        case receive := <-receiver:
            log.Println(receive)
        case <-time.After(time.Second): // 一定時間経過したら受信
            log.Println("timeout")
            return // 受信が終わればタイムアウト処理になる
        }
    }
}
