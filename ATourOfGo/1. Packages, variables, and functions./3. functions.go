package main

import(
    "fmt"
)

/*
変数名の後ろに型名を書く
*/
func add(x int, y int) int {
    return x + y
}

/*
関数の２つ以上の引数が同じ型である場合は、まとめることができる
*/
func add2(x, y int) int {
    return x + y
}

/*
複数の戻り値を返すことができる
*/
func swap(x, y string) (string, string) {
    return y, x
}

/*
戻り値となる変数に名前をつける（named return value）ことができる。
その場合、関数の最初で定義した変数名として扱われる
この方法を使うとreturnステートメントに何も書かずに戻すことができる（naked return）

※長い関数で使うと読みやすさに悪影響が出るので注意
*/
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(add(42, 13))
    fmt.Println(add2(42, 13))

    a, b := swap("hello", "world")
    fmt.Println(a, b) // => world hello

    fmt.Println(split(17))
}
