package main

import (
    "fmt"
)

/*
変数Tのポインタは *T型 で、ゼロ値は nil
var p *int

&オペレーターは、そのオペランドへのポインタを引き出す
i := 42
p = &i

*オペレーターは、ポインタの指す先の変数を示す
fmt.Println(*p) // ポインタpを通してiから値を読み出す
*p = 21 // ポインタpを通してiへ値を代入する
*/
func main() {
    i, j := 42, 2701

    p := &i
    fmt.Println(*p) // => 42
    *p = 21
    fmt.Println(i) // => 21

    p = &j
    *p = *p / 37
    fmt.Println(j) // => 73
}
