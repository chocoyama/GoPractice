package main

import (
    "log"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    for i:=0; i<3; i++ {
        wg.Add(1) // goroutine生成のたびにインクリメント
        go func(i int) {
            log.Println(i)
            wg.Done() // 終了時にデクリメント
        }(i)
    }
    wg.Wait() // ブロックして全てのDoneが終わったら次に進む
}
