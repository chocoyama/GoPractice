package main

import (
    "log"
)

/*
スケジュールを記述することで、処理を途中で止め、そこからリスタートすることができる
*/

func f(yield chan string) {
    yield <- "one"
    yield <- "two"
    yield <- "three"
}

func main() {
    co := make(chan string)
    go f(co)
    log.Println(<-co)
    log.Println(<-co)
    log.Println(<-co)
}
