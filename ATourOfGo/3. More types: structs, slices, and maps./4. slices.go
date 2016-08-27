package main

import (
    "fmt"
)

/*
[]Tは、T型の要素から成るスライス型を表現する
スライス型は値をストアしないので、元となる配列のアドレスを参照する
*/
func main() {
    primes := [6]int{2, 3, 5, 7, 11, 13}
    var s []int = primes[1:4] // 4番目の要素は含まない
    fmt.Println(s) // => [3, 5, 7]
    s = s[:2]
    fmt.Println(s) // => [3 5]
    s = s[1:]
    fmt.Println(s) // => [5]

    names := [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }
    fmt.Println(names) // => [John Paul George Ringo]
    a := names[0:2] // => [John Paul]
    b := names[1:3] // => [Paul George]
    fmt.Println(a, b)

    /*
    参照先ポインタの値が書きかわるので、同一アドレスを参照している全ての変数の値が書きかわる
    */
    b[0] = "XXX"
    fmt.Println(a, b) // => [John XXX] [XXX Geroge]
    fmt.Println(names) // => [John XXX Geroge Ringo]

    /*
    配列数を指定しない初期化はスライス型のリテラルとなる
    この場合、一旦配列を作成したのち、それを元にしたスライス型を生成して返す
    */
    q := []int{2, 3, 5, 7, 11, 13}
    fmt.Println(q)

    r := []bool{true, false, true, true, false ,true}
    fmt.Println(r)

    ss := []struct { //[]TのTの部分にstructのリテラルを入れることも可能
        i int
        b bool
    }{
        {2, true},
        {3, false},
        {5, true},
        {7, true},
        {11, false},
        {13, true},
    }
    fmt.Println(ss)
}
