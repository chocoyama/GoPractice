package main

import (
    "fmt"
    "time"
)

/*
goroutineは、Goのランタイムに管理される軽量なスレッド

go f(x, y, z)
と書けば、新しいgoroutineが実行される
f, x, y, zの評価は、実行元のgoroutineで実行され、fの実行は新しいgoroutineで実行される

goroutineは、同じアドレス空間で実行されるため、共有メモリへのアクセスは必ず同期する必要がある
syncパッケージは同期する際に役に立つ方法を提供しているが、別の方法(Channel)がある

チャネル（Channel）型は、チャネルオペレータの「<-」を用いて値の送受信ができる通り道
    ch := make(chan int)
    ch <- v // vをチャネルchへ送信する
    v := <-ch // chから受信した変数をvへ割り当てる

*/

func say(s string) {
    for i:=0; i<5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum
}

func main() {
    go say("world")
    say("hello")

    s := []int{7, 2, 8, -9, 4, 0}
    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // 2つのgoroutineが完了するまで処理がブロックされる
    fmt.Println(x, y, x+y)

    /*
    第二引数にバッファの長さを与えることができる
        バッファが詰まった時：チャネルへの送信をブロックする
        バッファが空の時：チャネルの受信をブロック
    */
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
