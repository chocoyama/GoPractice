package main

import(
    "fmt"
)

/*
varステートメントは変数を宣言する
関数の引数リストと同様に、複数の変数の最後に型を書くことで、変数のリストを宣言できる
*/
var a, b, c bool
var swift, python, java = true, false, "no!" // 初期化子を与えることができ、その場合型を省略できる
var ( // importステートメントと同様にまとめて宣言可能
    ToBe bool = false
    MaxInt uint64 = 1<<64 -1
    z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    var i, j int = 1, 2
    var n, m int
    k := 3 // varの代わりに「:=」で値を代入した上での変数宣言ができる（※ 関数の中のみ）

    fmt.Println(a, b, c) // => false false false
    fmt.Println(swift, python, java) // => true false no
    fmt.Println(i, j) // => 1 2
    fmt.Println(n, m) // => 0 0
    fmt.Println(k) // => 3
}
