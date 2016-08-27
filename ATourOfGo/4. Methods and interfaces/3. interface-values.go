package main

import (
    "fmt"
    "math"
)

type I interface {
    M()
}

type T struct {
    S string
}

func (t *T) M() {
    /*
    多くの言語はnilに対してメソッド呼び出しをするとぬるぽになるが、
    Goの場合は以下のようにメソッド内でハンドリングすることができる
    */
    if t == nil {
        fmt.Println("<nil>")
    } else {
        fmt.Println(t.S)
    }
}

type F float64

func (f F) M() {
    fmt.Println(f)
}

func main() {
    var i I

    i = &T{"Hello"}
    describe(i) // => (%{Hello}, *main.T) ※実装型として捉えられている
    i.M()

    i = F(math.Pi)
    describe(i) // => (3.141592653589793, main.F) ※実装型として捉えられている
    i.M()

    var t *T
    i = t
    describe(i) // => (<nil>, *main.T)
    i.M() // ぬるぽにならない!!（※メソッドが実装されていない || nilハンドリングされていない 場合はランタイムエラーになる可能性あり）

    /*
    var emptyi interface{} // 何も実装されていないinterface型の変数はAnyTypeとして使える
    emptyi = 42
    emptyi = "hello"
    */
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
