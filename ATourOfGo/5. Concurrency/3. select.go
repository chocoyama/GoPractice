package main

import (
    "fmt"
    "time"
)

/*
selectステートメントはgoroutineを複数の通信操作で待たせる
*/

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        /*
        複数あるcaseのいずれかが準備できるようになるまでブロックし、準備ができたcaseを実行する
        複数のcaseの準備ができている場合、caseはランダムに選択される。
        */
        select {
        case c <- x:
            x, y = y, x + y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func tickboom() {
    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)
    for {
        select {
        case <-tick:
            fmt.Println("tick.")
        case <-boom:
            fmt.Println("BOOM!")
            return
        default: // どのcaseも準備できていない場合は、defaultが実行される
            fmt.Println("    .")
            time.Sleep(50 * time.Millisecond)
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i:=0; i<10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)

    tickboom()
}
