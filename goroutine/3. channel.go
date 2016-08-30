package main

import (
    "log"
)

/*
channelとは
- goroutine間でのメッセージパッシングをするためのもの
- メッセージの型を指定できる
- first class valueであり、引数や戻り値にも使える
- send/receiveでブロックする
- bufferで、一度に扱えるメッセージ量を指定できる

make()でインスタンスを生成
送信: channel<-value
受信: <-channel
*/

func f(ch chan bool) {
    ch <- true
}

func main() {
    ch := make(chan bool)
    go f(ch)
    log.Println(<-ch) // ここでデータが来るまでブロックされる
}
