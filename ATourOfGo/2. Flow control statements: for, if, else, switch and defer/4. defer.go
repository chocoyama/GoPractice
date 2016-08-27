package main

import (
    "fmt"
)

func main() {
    /*
    deferへ渡した関数の実行を、呼び出し元の関数の終わりまで遅延する
    渡した関数の引数はすぐに評価されるが、関数自体は呼び出し元の関数がreturnするまで実行されない

    deferへ渡した関数が複数ある場合は、その呼び出しはスタックされる
    呼び出し元の関数がreturnするとき、deferへ渡した関数がLIFOの順番で実行される
    */
    defer fmt.Println("world")
    fmt.Println("hello") // => hello ... world

    fmt.Println("counting")
    for i:=0; i<10; i++ {
        defer fmt.Println(i) // 9, 8, ..., 1, 0
    }
    fmt.Println("done") // counting done 9, 8, ..., 1, 0
}
