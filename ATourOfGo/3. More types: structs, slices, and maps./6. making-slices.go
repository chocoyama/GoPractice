package main

import (
    "fmt"
)

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

/*
第一引数：型
第二引数：length
第三引数：capacity（省略するとlengthと同じ値になる）
*/
func main() {
    a := make([]int, 5, 5)
    printSlice("a", a) // => a len=5 cap=5 [0 0 0 0 0]

    b := make([]int, 5)
    printSlice("b", b) // => b len=5 cap=5 [0 0 0 0 0]

    c := make([]int, 0, 5)
    printSlice("c", c) // => c len=0 cap=5 []

    d := c[:2]
    printSlice("d", d) // => d len=2 cap=5 [0 0]

    e := d[2:5]
    printSlice("e", e) // => e len=3 cap=3 [0 0 0]
}
