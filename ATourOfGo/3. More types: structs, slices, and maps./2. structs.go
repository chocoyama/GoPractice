package main

import (
    "fmt"
)

/*
structはフィールドの集まり
*/
type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v) // => {4 2}
    fmt.Println(v.X) // => 4

    /*
    ポインタを通したアクセスも可能
    */
    p := &v
    p.X = 1e9
    fmt.Println(v) // => {1000000000 2}

    /*
    フィールドに初期値を与えない場合、ゼロ値が与えられる
    */
    v2 := Vertex{X: 1}
    v3 := Vertex{}
    p = &Vertex{1, 2}
    fmt.Println(v2, v3, p) // => {1 0} {0 0} &{1 2}
}
