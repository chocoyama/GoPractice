package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
    /*
    forのように条件の前に、評価するための簡単なステートメントを書くことができる
    ここで宣言された変数はifのスコープ内でだけ有効
    */
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        // ifステートメントで宣言された変数はelseブロック内でも使うことができる
        fmt.Printf("%g >= %g\n", v, lim)
        return lim
    }
}

func main() {
    fmt.Println(sqrt(2), sqrt(-4)) // => 1.4142135623730951 2i

    fmt.Println(
        pow(3, 2, 10), // => 9
        pow(3, 3, 20), // => 20
    )
}
