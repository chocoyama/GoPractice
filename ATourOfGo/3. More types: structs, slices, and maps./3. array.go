package main

import (
    "fmt"
)

/*
[n]T型は、型Tのn個の変数の配列を表す
var a [10]int

Goでは配列の長さは、型の一部分であるので、配列のサイズを変更することはできない
また、配列が格納されている変数は、C言語のように先頭配列要素へのポインタではなく、配列全体を表している
*/
func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1]) // => Hello World
    fmt.Println(a) // => [Hello World]

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes) // => [2 3 5 7 11 13]

}
