package main

import (
    "fmt"
    "sync"
    "time"
)

/*
排他制御
このデータ構造を指す一般的な名前はMutex

Goの標準ライブラリは、排他制御をsync.Mutexと次の二つのメソッドで提供する
Lock
Unlock
*/

type SafeCounter struct {
    v map[string]int
    mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
    /*
    LockとUnlockで囲むことで排他制御で実行するコードを定義する
    */
    c.mux.Lock()
    c.v[key]++
    c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
    c.mux.Lock()
    defer c.mux.Unlock() // mutexがUnlockされることを保証する
    return c.v[key]
}

func main() {
    c := SafeCounter{v: make(map[string]int)}
    for i:=0; i<1000; i++ {
        go c.Inc("somekey")
    }

    time.Sleep(time.Second)
    fmt.Println(c.Value("somekey"))
}
