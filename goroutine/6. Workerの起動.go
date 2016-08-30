package main

import (
    "fmt"
    "log"
)

/*
「<-chan」を返り値の型にしているので、呼び出し先からは書き込みできなくなっている
*/
func worker(msg string) <-chan string {
    receiver := make(chan string)
    for i:=0; i<3; i++ {
        go func(i int) {
            msg := fmt.Sprintf("%d %s done", i, msg)
            receiver <- msg // goroutineの中で送信をしている
        }(i)
    }
    return receiver
}

func main() {
    receiver := worker("job")
    for i:=0; i<3; i++ {
        log.Println(<-receiver)
    }
}
