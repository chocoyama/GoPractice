package main

import (
    "log"
)

func main() {
    fin := make(chan bool)
    go func() {
        log.Println("worker working...")
        fin <- false
    }()

    /*
    channelは、送受信が完了するまでブロックする
    そのため、下記のようにchannelを受信していると、goroutineが終了するまで処理が終了しない
    */
    <-fin
}
