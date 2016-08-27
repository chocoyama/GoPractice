package main

import (
    "fmt"
)

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i:=0; i<n; i++ {
        c <- x
        x, y = y, x+y
    }
    /*
    これ以上送信する値がない場合を示すには、チャネルをcloseする
    受信式に2つ目のパラメータを割りあてることで、そのチャネルがcloseされているかどうか確認できる
        v, ok := <-ch
        受信する値がなく、チャネルが閉じている場合、okの変数はfalseになる
    */
    close(c)
}

func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    for i := range c { // チャネルが閉じられるまで、チャネルから値を繰り返し受信し続ける
        fmt.Println(i)
    }
}
